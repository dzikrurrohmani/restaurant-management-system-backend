package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer []*model.Customer) error
	FindBy(by map[string]interface{}) ([]model.Customer, error)
	FindAll() ([]model.Customer, error)
	UpdateBy(customer *model.Customer, by map[string]interface{}) error
	Delete(customer *model.Customer) error
	UpdateAssociation(assocModel *model.Customer, assocName string, assocNewValue interface{}) error
}
type customerRepository struct {
	db *gorm.DB
}

func (m *customerRepository) Create(customer []*model.Customer) error {
	result := m.db.Create(customer)
	return result.Error
}

func (m *customerRepository) FindBy(by map[string]interface{}) ([]model.Customer, error) {
	var customers []model.Customer
	result := m.db.Preload("Discounts").Where(by).Find(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return customers, nil
}

func (m *customerRepository) FindAll() ([]model.Customer, error) {
	var customers []model.Customer
	result := m.db.Preload("Discounts").Find(&customers)
	if err := result.Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (m *customerRepository) UpdateBy(customer *model.Customer, by map[string]interface{}) error {
	result := m.db.Model(customer).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (m *customerRepository) Delete(customer *model.Customer) error {
	result := m.db.Delete(&model.Customer{}, customer).Error
	return result
}

func (c *customerRepository) UpdateAssociation(assocModel *model.Customer, assocName string, assocNewValue interface{}) error {
	err := c.db.Model(assocModel).Association(assocName).Replace(assocNewValue)
	if err != nil {
		return err
	}
	return nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
