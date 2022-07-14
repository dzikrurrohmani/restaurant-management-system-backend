package discount_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type UpdateDiscountUseCase interface {
	UpdateDiscount(discount *model.Discount, by map[string]interface{}) error
}

type updateDiscountUseCase struct {
	discountRepo repository.DiscountRepository
}

func (m *updateDiscountUseCase) UpdateDiscount(discount *model.Discount, by map[string]interface{}) error {
	return m.discountRepo.UpdateBy(discount, by)
}

func NewUpdateDiscountUseCase(discountRepo repository.DiscountRepository) UpdateDiscountUseCase {
	usecase := new(updateDiscountUseCase)
	usecase.discountRepo = discountRepo
	return usecase
}
