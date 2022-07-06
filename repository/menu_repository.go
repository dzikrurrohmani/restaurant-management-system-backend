package repository

import (
	"errors"
	"livecode-wmb-rest-api/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(menu []*model.Menu) error
	FindBy(by map[string]interface{}) ([]model.Menu, error)
	FindAll() ([]model.Menu, error)
	UpdateBy(menu *model.Menu, by map[string]interface{}) error
	Delete(menu *model.Menu) error
}
type menuRepository struct {
	db *gorm.DB
}

func (m *menuRepository) Create(menu []*model.Menu) error {
	result := m.db.Create(menu)
	return result.Error
}

func (m *menuRepository) FindBy(by map[string]interface{}) ([]model.Menu, error) {
	var menus []model.Menu
	result := m.db.Where(by).Find(&menus)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return menus, nil
}

func (m *menuRepository) FindAll() ([]model.Menu, error) {
	var menus []model.Menu
	result := m.db.Find(&menus)
	if err := result.Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (m *menuRepository) UpdateBy(menu *model.Menu, by map[string]interface{}) error {
	result := m.db.Model(menu).Updates(by)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
func (m *menuRepository) Delete(menu *model.Menu) error {
	result := m.db.Delete(&model.Menu{}, menu).Error
	return result
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	repo := new(menuRepository)
	repo.db = db
	return repo
}
