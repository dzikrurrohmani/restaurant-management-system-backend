package discount_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type DeleteDiscountUseCase interface {
	DeleteDiscount(discount *model.Discount) error
}

type deleteDiscountUseCase struct {
	discountRepo repository.DiscountRepository
}

func (m *deleteDiscountUseCase) DeleteDiscount(discount *model.Discount) error {
	return m.discountRepo.Delete(discount)
}

func NewDeleteDiscountUseCase(discountRepo repository.DiscountRepository) DeleteDiscountUseCase {
	usecase := new(deleteDiscountUseCase)
	usecase.discountRepo = discountRepo
	return usecase
}
