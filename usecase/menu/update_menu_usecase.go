package menu_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type UpdateMenuUseCase interface {
	UpdateMenu(menu *model.Menu, by map[string]interface{}) error
}

type updateMenuUseCase struct {
	menuRepo repository.MenuRepository
}

func (m *updateMenuUseCase) UpdateMenu(menu *model.Menu, by map[string]interface{}) error {
	return m.menuRepo.UpdateBy(menu, by)
}

func NewUpdateMenuUseCase(menuRepo repository.MenuRepository) UpdateMenuUseCase {
	usecase := new(updateMenuUseCase)
	usecase.menuRepo = menuRepo
	return usecase
}
