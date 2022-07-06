package discount_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
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
