package repository

import (
	"errors"
	"livecode-wmb-2/model"

	"gorm.io/gorm"
)

type MenuPriceRepository interface {
	Create(menuPrice []*model.MenuPrice) error
	FindBy(by map[string]interface{}) ([]model.MenuPrice, error)
	FindAll() ([]model.MenuPrice, error)
	UpdateBy(menuPrice *model.MenuPrice, by map[string]interface{}) error
	Delete(menuPrice *model.MenuPrice) error
}
type menuPriceRepository struct {
	db *gorm.DB
}

func (m *menuPriceRepository) Create(menuPrice []*model.MenuPrice) error {
	result := m.db.Create(menuPrice)
	return result.Error
}

func (m *menuPriceRepository) FindBy(by map[string]interface{}) ([]model.MenuPrice, error) {
	var menuPrices []model.MenuPrice
	result := m.db.Where(by).Find(&menuPrices)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return menuPrices, nil
}

func (m *menuPriceRepository) FindAll() ([]model.MenuPrice, error) {
	var menuPrices []model.MenuPrice
	result := m.db.Find(&menuPrices)
	if err := result.Error; err != nil {
		return nil, err
	}
	return menuPrices, nil
}

func (m *menuPriceRepository) UpdateBy(menuPrice *model.MenuPrice, by map[string]interface{}) error {
	result := m.db.Model(menuPrice).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
func (m *menuPriceRepository) Delete(menuPrice *model.MenuPrice) error {
	result := m.db.Delete(&model.MenuPrice{}, menuPrice).Error
	return result
}

func NewMenuPriceRepository(db *gorm.DB) MenuPriceRepository {
	repo := new(menuPriceRepository)
	repo.db = db
	return repo
}
