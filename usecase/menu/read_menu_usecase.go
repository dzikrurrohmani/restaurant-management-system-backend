package menu_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type ReadMenuUseCase interface {
	ReadAllMenu() ([]model.Menu, error)
	ReadMenuById(id uint) (model.Menu, error)
	ReadMenuBy(by map[string]interface{}) ([]model.Menu, error)
}

type readMenuUseCase struct {
	menuRepo repository.MenuRepository
}

func (m *readMenuUseCase) ReadAllMenu() ([]model.Menu, error) {
	return m.menuRepo.FindAll()
}

func (m *readMenuUseCase) ReadMenuById(id uint) (model.Menu, error) {
	menuSlice, err := m.menuRepo.FindBy(map[string]interface{}{"id": id})
	if len(menuSlice) == 0 {
		return model.Menu{}, err
	}
	return menuSlice[0], err
}

func (m *readMenuUseCase) ReadMenuBy(by map[string]interface{}) ([]model.Menu, error) {
	return m.menuRepo.FindBy(by)
}

func NewReadMenuUseCase(menuRepo repository.MenuRepository) ReadMenuUseCase {
	usecase := new(readMenuUseCase)
	usecase.menuRepo = menuRepo
	return usecase
}
