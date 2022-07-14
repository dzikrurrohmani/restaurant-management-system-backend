package transtype_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type UpdateTransTypeUseCase interface {
	UpdateTransType(transType *model.TransType, by map[string]interface{}) error
}

type updateTransTypeUseCase struct {
	transTypeRepo repository.TransTypeRepository
}

func (m *updateTransTypeUseCase) UpdateTransType(transType *model.TransType, by map[string]interface{}) error {
	return m.transTypeRepo.UpdateBy(transType, by)
}

func NewUpdateTransTypeUseCase(transTypeRepo repository.TransTypeRepository) UpdateTransTypeUseCase {
	usecase := new(updateTransTypeUseCase)
	usecase.transTypeRepo = transTypeRepo
	return usecase
}
