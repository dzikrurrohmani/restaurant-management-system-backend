package table_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type CreateTableUseCase interface {
	CreateTable(table []*model.Table) error
}

type createTableUseCase struct {
	tableRepo repository.TableRepository
}

func (m *createTableUseCase) CreateTable(table []*model.Table) error {
	return m.tableRepo.Create(table)
}

func NewCreateTableUseCase(tableRepo repository.TableRepository) CreateTableUseCase {
	usecase := new(createTableUseCase)
	usecase.tableRepo = tableRepo
	return usecase
}
