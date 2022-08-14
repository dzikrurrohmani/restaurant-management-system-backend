package repository

import (
	"errors"
	"fmt"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type TableRepository interface {
	Create(table []*model.Table) error
	FindBy(by map[string]interface{}) ([]model.Table, error)
	FindAll() ([]model.Table, error)
	UpdateBy(table *model.Table, by map[string]interface{}) error
	Delete(table *model.Table) error
}
type tableRepository struct {
	db *gorm.DB
}

func (m *tableRepository) Create(table []*model.Table) error {
	fmt.Println(table[0].IsAvailable)
	result := m.db.Create(table)
	return result.Error
}

func (m *tableRepository) FindBy(by map[string]interface{}) ([]model.Table, error) {
	var tables []model.Table
	result := m.db.Where(by).Find(&tables)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return tables, nil
}

func (m *tableRepository) FindAll() ([]model.Table, error) {
	var tables []model.Table
	result := m.db.Find(&tables)
	if err := result.Error; err != nil {
		return nil, err
	}
	return tables, nil
}

func (m *tableRepository) UpdateBy(table *model.Table, by map[string]interface{}) error {
	result := m.db.Model(table).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
func (m *tableRepository) Delete(table *model.Table) error {
	result := m.db.Delete(&model.Table{}, table).Error
	return result
}

func NewTableRepository(db *gorm.DB) TableRepository {
	repo := new(tableRepository)
	repo.db = db
	return repo
}
