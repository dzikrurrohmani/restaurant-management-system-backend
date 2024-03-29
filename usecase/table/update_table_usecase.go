package table_usecase

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
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
