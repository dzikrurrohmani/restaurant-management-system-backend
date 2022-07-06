package menu_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type CreateMenuUseCase interface {
	CreateMenu(menu []*model.Menu) error
}

type createMenuUseCase struct {
	menuRepo repository.MenuRepository
}

func (m *createMenuUseCase) CreateMenu(menu []*model.Menu) error {
	return m.menuRepo.Create(menu)
}

func NewCreateMenuUseCase(menuRepo repository.MenuRepository) CreateMenuUseCase {
	usecase := new(createMenuUseCase)
	usecase.menuRepo = menuRepo
	return usecase
}
