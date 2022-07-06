package table_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type UpdateTableUseCase interface {
	UpdateTable(table *model.Table, by map[string]interface{}) error
}

type updateTableUseCase struct {
	tableRepo repository.TableRepository
}

func (m *updateTableUseCase) UpdateTable(table *model.Table, by map[string]interface{}) error {
	return m.tableRepo.UpdateBy(table, by)
}

func NewUpdateTableUseCase(tableRepo repository.TableRepository) UpdateTableUseCase {
	usecase := new(updateTableUseCase)
	usecase.tableRepo = tableRepo
	return usecase
}
