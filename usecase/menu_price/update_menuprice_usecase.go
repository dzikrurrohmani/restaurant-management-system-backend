package menuprice_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type UpdateMenuPriceUseCase interface {
	UpdateMenuPrice(menuPrice *model.MenuPrice, by map[string]interface{}) error
}

type updateMenuPriceUseCase struct {
	menuPriceRepo repository.MenuPriceRepository
}

func (m *updateMenuPriceUseCase) UpdateMenuPrice(menuPrice *model.MenuPrice, by map[string]interface{}) error {
	return m.menuPriceRepo.UpdateBy(menuPrice, by)
}

func NewUpdateMenuPriceUseCase(menuPriceRepo repository.MenuPriceRepository) UpdateMenuPriceUseCase {
	usecase := new(updateMenuPriceUseCase)
	usecase.menuPriceRepo = menuPriceRepo
	return usecase
}
