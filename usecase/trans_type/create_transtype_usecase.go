package transtype_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type CreateTransTypeUseCase interface {
	CreateTransType(transType []*model.TransType) error
}

type createTransTypeUseCase struct {
	transTypeRepo repository.TransTypeRepository
}

func (m *createTransTypeUseCase) CreateTransType(transType []*model.TransType) error {
	return m.transTypeRepo.Create(transType)
}

func NewCreateTransTypeUseCase(transTypeRepo repository.TransTypeRepository) CreateTransTypeUseCase {
	usecase := new(createTransTypeUseCase)
	usecase.transTypeRepo = transTypeRepo
	return usecase
}
