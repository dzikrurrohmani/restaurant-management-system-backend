package menuprice_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type DeleteMenuPriceUseCase interface {
	DeleteMenuPrice(menu *model.MenuPrice) error
}

type deleteMenuPriceUseCase struct {
	menuPriceRepo repository.MenuPriceRepository
}

func (m *deleteMenuPriceUseCase) DeleteMenuPrice(menu *model.MenuPrice) error {
	return m.menuPriceRepo.Delete(menu)
}

func NewDeleteMenuPriceUseCase(menuPriceRepo repository.MenuPriceRepository) DeleteMenuPriceUseCase {
	usecase := new(deleteMenuPriceUseCase)
	usecase.menuPriceRepo = menuPriceRepo
	return usecase
}
