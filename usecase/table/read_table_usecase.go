package table_usecase

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type ReadTableUseCase interface {
	ReadAllTable() ([]model.Table, error)
	ReadTableById(id uint) (model.Table, error)
	ReadTableBy(by map[string]interface{}) ([]model.Table, error)
}

type readTableUseCase struct {
	tableRepo repository.TableRepository
}

func (m *readTableUseCase) ReadAllTable() ([]model.Table, error) {
	return m.tableRepo.FindAll()
}

func (m *readTableUseCase) ReadTableById(id uint) (model.Table, error) {
	tableSlice, err := m.tableRepo.FindBy(map[string]interface{}{"id": id})
	if len(tableSlice) == 0 {
		return model.Table{}, err
	}
	return tableSlice[0], err
}

func (m *readTableUseCase) ReadTableBy(by map[string]interface{}) ([]model.Table, error) {
	return m.tableRepo.FindBy(by)
}

func NewReadTableUseCase(tableRepo repository.TableRepository) ReadTableUseCase {
	usecase := new(readTableUseCase)
	usecase.tableRepo = tableRepo
	return usecase
}
