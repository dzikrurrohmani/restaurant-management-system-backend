package menuprice_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type CreateMenuPriceUseCase interface {
	CreateMenuPrice(menuPrice []*model.MenuPrice) error
}

type createMenuPriceUseCase struct {
	menuPriceRepo repository.MenuPriceRepository
}

func (m *createMenuPriceUseCase) CreateMenuPrice(menuPrice []*model.MenuPrice) error {
	return m.menuPriceRepo.Create(menuPrice)
}

func NewCreateMenuPriceUseCase(menuPriceRepo repository.MenuPriceRepository) CreateMenuPriceUseCase {
	usecase := new(createMenuPriceUseCase)
	usecase.menuPriceRepo = menuPriceRepo
	return usecase
}
