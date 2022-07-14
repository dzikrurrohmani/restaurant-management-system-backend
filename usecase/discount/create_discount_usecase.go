package discount_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type CreateDiscountUseCase interface {
	CreateDiscount(dicsount []*model.Discount) error
}

type createDiscountUseCase struct {
	dicsountRepo repository.DiscountRepository
}

func (m *createDiscountUseCase) CreateDiscount(dicsount []*model.Discount) error {
	return m.dicsountRepo.Create(dicsount)
}

func NewCreateDiscountUseCase(dicsountRepo repository.DiscountRepository) CreateDiscountUseCase {
	usecase := new(createDiscountUseCase)
	usecase.dicsountRepo = dicsountRepo
	return usecase
}
