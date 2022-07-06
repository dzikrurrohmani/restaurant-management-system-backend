package billdetail_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type GetIncomeUseCase interface {
	GenDailyIncome() ([]model.MenuPrice, error)
}

type readMenuPriceUseCase struct {
	menuPriceRepo repository.MenuPriceRepository
}

func (m *readMenuPriceUseCase) ReadAllMenuPrice() ([]model.MenuPrice, error) {
	return m.menuPriceRepo.FindAll()
}

func (m *readMenuPriceUseCase) GetIncomeById(id uint) (model.MenuPrice, error) {
	menuPriceSlice, err := m.menuPriceRepo.FindBy(map[string]interface{}{"id": id})
	if len(menuPriceSlice) == 0 {
		return model.MenuPrice{}, err
	}
	return menuPriceSlice[0], err
}

func (m *readMenuPriceUseCase) GetIncomeBy(by map[string]interface{}) ([]model.MenuPrice, error) {
	return m.menuPriceRepo.FindBy(by)
}

func NewGetIncomeUseCase(menuPriceRepo repository.MenuPriceRepository) GetIncomeUseCase {
	usecase := new(readMenuPriceUseCase)
	usecase.menuPriceRepo = menuPriceRepo
	return usecase
}
