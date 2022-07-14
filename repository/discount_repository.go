package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	Create(discount []*model.Discount) error
	FindBy(by map[string]interface{}) ([]model.Discount, error)
	FindAll() ([]model.Discount, error)
	UpdateBy(discount *model.Discount, by map[string]interface{}) error
	Delete(discount *model.Discount) error
}
type discountRepository struct {
	db *gorm.DB
}

func (m *discountRepository) Create(discount []*model.Discount) error {
	result := m.db.Create(discount)
	return result.Error
}

func (m *discountRepository) FindBy(by map[string]interface{}) ([]model.Discount, error) {
	var discounts []model.Discount
	result := m.db.Where(by).Find(&discounts)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return discounts, nil
}

func (m *discountRepository) FindAll() ([]model.Discount, error) {
	var discounts []model.Discount
	result := m.db.Find(&discounts)
	if err := result.Error; err != nil {
		return nil, err
	}
	return discounts, nil
}

func (m *discountRepository) UpdateBy(discount *model.Discount, by map[string]interface{}) error {
	result := m.db.Model(discount).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
func (m *discountRepository) Delete(discount *model.Discount) error {
	result := m.db.Delete(&model.Discount{}, discount).Error
	return result
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	repo := new(discountRepository)
	repo.db = db
	return repo
}
