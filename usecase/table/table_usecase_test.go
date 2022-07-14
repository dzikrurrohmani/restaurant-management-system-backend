package table_usecase

import (
	"errors"
	"testing"

	"livecode-wmb-2/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyTables = []model.Table{
	{
		ID:               1,
		TableDescription: "T001",
		IsAvailable:      true,
		Bills:            nil,
	},
	{
		ID:               2,
		TableDescription: "T002",
		IsAvailable:      true,
		Bills:            nil,
	},
}

type tableRepoMock struct {
	mock.Mock
}

type TableUseCaseTestSuite struct {
	suite.Suite
	tableRepoMock *tableRepoMock
	// tableRepoMock repository.TableRepository
}

func (r *tableRepoMock) Create(table []*model.Table) error {
	args := r.Called(table)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *tableRepoMock) FindBy(by map[string]interface{}) ([]model.Table, error) {
	args := r.Called(by)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Table), nil
}
func (r *tableRepoMock) FindAll() ([]model.Table, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Table), nil
}
func (r *tableRepoMock) UpdateBy(table *model.Table, by map[string]interface{}) error {
	args := r.Called(table, by)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *tableRepoMock) Delete(table *model.Table) error {
	args := r.Called(table)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite *TableUseCaseTestSuite) SetupTest() {
	suite.tableRepoMock = new(tableRepoMock)
}

func (suite *TableUseCaseTestSuite) TestCreateTable() {
	dummyTable := dummyTables[0]
	suite.tableRepoMock.On("Create", []*model.Table{&dummyTable}).Return(nil) // expected, Get(0) : dummyTable, Get(1) : nil
	tableUseCase := NewCreateTableUseCase(suite.tableRepoMock)
	err := tableUseCase.CreateTable([]*model.Table{&dummyTable}) // actual
	assert.Nil(suite.T(), err)
}

func (suite *TableUseCaseTestSuite) TestDeleteTable() {
	dummyTable := dummyTables[0]
	suite.tableRepoMock.On("Delete", &dummyTable).Return(nil) // expected, Get(0) : dummyTable, Get(1) : nil
	tableUseCase := NewDeleteTableUseCase(suite.tableRepoMock)
	err := tableUseCase.DeleteTable(&dummyTable) // actual
	assert.Nil(suite.T(), err)
}

func (suite *TableUseCaseTestSuite) TestUpdateTable() {
	var by = map[string]interface{}{
		"Id": 1,
	}
	dummyTable := dummyTables[0]
	suite.tableRepoMock.On("UpdateBy", &dummyTable, by).Return(nil) // expected, Get(0) : dummyTable, Get(1) : nil
	tableUseCase := NewUpdateTableUseCase(suite.tableRepoMock)
	err := tableUseCase.UpdateTable(&dummyTable, by) // actual
	assert.Nil(suite.T(), err)
}

func (suite *TableUseCaseTestSuite) TestTableFindById_Success() {
	dummyTable := dummyTables[0]
	// Mengacu ke nama method
	suite.tableRepoMock.On("FindBy", map[string]interface{}{"id": dummyTable.ID}).Return([]model.Table{dummyTable}, nil)
	tableUseCase := NewReadTableUseCase(suite.tableRepoMock)
	table, err := tableUseCase.ReadTableById(dummyTable.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTable.ID, table.ID)
}

func (suite *TableUseCaseTestSuite) TestTableFindById_Failed() {
	dummyTable := dummyTables[0]
	// Mengacu ke nama method
	suite.tableRepoMock.On("FindBy", map[string]interface{}{"id": dummyTable.ID}).Return(nil, errors.New("failed"))
	tableUseCase := NewReadTableUseCase(suite.tableRepoMock)
	table, err := tableUseCase.ReadTableById(dummyTable.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Equal(suite.T(), model.Table{}, table)
}

func (suite *TableUseCaseTestSuite) TestTableFindBy() {
	dummyTable := dummyTables[0]
	var by = map[string]interface{}{"id": dummyTable.ID}
	// Mengacu ke nama method
	suite.tableRepoMock.On("FindBy", by).Return([]model.Table{dummyTable}, nil)
	tableUseCase := NewReadTableUseCase(suite.tableRepoMock)
	table, err := tableUseCase.ReadTableBy(by)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTable.ID, table[0].ID)
}

func (suite *TableUseCaseTestSuite) TestTableReadeAll() {
	// Mengacu ke nama method
	suite.tableRepoMock.On("FindAll").Return(dummyTables, nil)
	tableUseCase := NewReadTableUseCase(suite.tableRepoMock)
	tables, err := tableUseCase.ReadAllTable()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTables, tables)
}

func TestTablePriceUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TableUseCaseTestSuite))
}
