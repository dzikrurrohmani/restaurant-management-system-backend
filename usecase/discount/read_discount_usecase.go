package discount_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type ReadDiscountUseCase interface {
	ReadAllDiscount() ([]model.Discount, error)
	ReadDiscountById(id uint) (model.Discount, error)
	ReadDiscountBy(by map[string]interface{}) ([]model.Discount, error)
}

type readDiscountUseCase struct {
	discountRepo repository.DiscountRepository
}

func (m *readDiscountUseCase) ReadAllDiscount() ([]model.Discount, error) {
	return m.discountRepo.FindAll()
}

func (m *readDiscountUseCase) ReadDiscountById(id uint) (model.Discount, error) {
	discountSlice, err := m.discountRepo.FindBy(map[string]interface{}{"id": id})
	if len(discountSlice) == 0 {
		return model.Discount{}, err
	}
	return discountSlice[0], err
}

func (m *readDiscountUseCase) ReadDiscountBy(by map[string]interface{}) ([]model.Discount, error) {
	return m.discountRepo.FindBy(by)
}

func NewReadDiscountUseCase(discountRepo repository.DiscountRepository) ReadDiscountUseCase {
	usecase := new(readDiscountUseCase)
	usecase.discountRepo = discountRepo
	return usecase
}
