package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type TransTypeRepository interface {
	Create(transType []*model.TransType) error
	FindBy(by map[string]interface{}) ([]model.TransType, error)
	FindAll() ([]model.TransType, error)
	UpdateBy(transType *model.TransType, by map[string]interface{}) error
	Delete(transType *model.TransType) error
}
type transTypeRepository struct {
	db *gorm.DB
}

func (m *transTypeRepository) Create(transType []*model.TransType) error {
	result := m.db.Create(transType)
	return result.Error
}

func (m *transTypeRepository) FindBy(by map[string]interface{}) ([]model.TransType, error) {
	var transTypes []model.TransType
	result := m.db.Where(by).Find(&transTypes)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return transTypes, nil
}

func (m *transTypeRepository) FindAll() ([]model.TransType, error) {
	var transTypes []model.TransType
	result := m.db.Find(&transTypes)
	if err := result.Error; err != nil {
		return nil, err
	}
	return transTypes, nil
}

func (m *transTypeRepository) UpdateBy(transType *model.TransType, by map[string]interface{}) error {
	result := m.db.Model(transType).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
func (m *transTypeRepository) Delete(transType *model.TransType) error {
	result := m.db.Delete(&model.TransType{}, transType).Error
	return result
}

func NewTransTypeRepository(db *gorm.DB) TransTypeRepository {
	repo := new(transTypeRepository)
	repo.db = db
	return repo
}
