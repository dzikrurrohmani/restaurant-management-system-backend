package transtype_usecase

import (
	"errors"
	"testing"

	"livecode-wmb-2/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyTransTypes = []model.TransType{
	{
		ID:          "EI",
		Description: "Eat In",
		Bills:       nil,
	},
	{
		ID:          "TA",
		Description: "Take Away",
		Bills:       nil,
	},
}

type transTypeRepoMock struct {
	mock.Mock
}

type TransTypeUseCaseTestSuite struct {
	suite.Suite
	transTypeRepoMock *transTypeRepoMock
	// transTypeRepoMock repository.TransTypeRepository
}

func (r *transTypeRepoMock) Create(transType []*model.TransType) error {
	args := r.Called(transType)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *transTypeRepoMock) FindBy(by map[string]interface{}) ([]model.TransType, error) {
	args := r.Called(by)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TransType), nil
}
func (r *transTypeRepoMock) FindAll() ([]model.TransType, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.TransType), nil
}
func (r *transTypeRepoMock) UpdateBy(transType *model.TransType, by map[string]interface{}) error {
	args := r.Called(transType, by)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *transTypeRepoMock) Delete(transType *model.TransType) error {
	args := r.Called(transType)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite *TransTypeUseCaseTestSuite) SetupTest() {
	suite.transTypeRepoMock = new(transTypeRepoMock)
}

func (suite *TransTypeUseCaseTestSuite) TestCreateTransType() {
	dummyTransType := dummyTransTypes[0]
	suite.transTypeRepoMock.On("Create", []*model.TransType{&dummyTransType}).Return(nil) // expected, Get(0) : dummyTransType, Get(1) : nil
	transTypeUseCase := NewCreateTransTypeUseCase(suite.transTypeRepoMock)
	err := transTypeUseCase.CreateTransType([]*model.TransType{&dummyTransType}) // actual
	assert.Nil(suite.T(), err)
}

func (suite *TransTypeUseCaseTestSuite) TestDeleteTransType() {
	dummyTransType := dummyTransTypes[0]
	suite.transTypeRepoMock.On("Delete", &dummyTransType).Return(nil) // expected, Get(0) : dummyTransType, Get(1) : nil
	transTypeUseCase := NewDeleteTransTypeUseCase(suite.transTypeRepoMock)
	err := transTypeUseCase.DeleteTransType(&dummyTransType) // actual
	assert.Nil(suite.T(), err)
}

func (suite *TransTypeUseCaseTestSuite) TestUpdateTransType() {
	var by = map[string]interface{}{
		"Id": 1,
	}
	dummyTransType := dummyTransTypes[0]
	suite.transTypeRepoMock.On("UpdateBy", &dummyTransType, by).Return(nil) // expected, Get(0) : dummyTransType, Get(1) : nil
	transTypeUseCase := NewUpdateTransTypeUseCase(suite.transTypeRepoMock)
	err := transTypeUseCase.UpdateTransType(&dummyTransType, by) // actual
	assert.Nil(suite.T(), err)
}

func (suite *TransTypeUseCaseTestSuite) TestTransTypeFindById_Success() {
	dummyTransType := dummyTransTypes[0]
	// Mengacu ke nama method
	suite.transTypeRepoMock.On("FindBy", map[string]interface{}{"id": dummyTransType.ID}).Return([]model.TransType{dummyTransType}, nil)
	transTypeUseCase := NewReadTransTypeUseCase(suite.transTypeRepoMock)
	transType, err := transTypeUseCase.ReadTransTypeById(dummyTransType.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTransType.ID, transType.ID)
}

func (suite *TransTypeUseCaseTestSuite) TestTransTypeFindById_Failed() {
	dummyTransType := dummyTransTypes[0]
	// Mengacu ke nama method
	suite.transTypeRepoMock.On("FindBy", map[string]interface{}{"id": dummyTransType.ID}).Return(nil, errors.New("failed"))
	transTypeUseCase := NewReadTransTypeUseCase(suite.transTypeRepoMock)
	transType, err := transTypeUseCase.ReadTransTypeById(dummyTransType.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Equal(suite.T(), model.TransType{}, transType)
}

func (suite *TransTypeUseCaseTestSuite) TestTransTypeFindBy() {
	dummyTransType := dummyTransTypes[0]
	var by = map[string]interface{}{"id": dummyTransType.ID}
	// Mengacu ke nama method
	suite.transTypeRepoMock.On("FindBy", by).Return([]model.TransType{dummyTransType}, nil)
	transTypeUseCase := NewReadTransTypeUseCase(suite.transTypeRepoMock)
	transType, err := transTypeUseCase.ReadTransTypeBy(by)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTransType.ID, transType[0].ID)
}

func (suite *TransTypeUseCaseTestSuite) TestTransTypeReadeAll() {
	// Mengacu ke nama method
	suite.transTypeRepoMock.On("FindAll").Return(dummyTransTypes, nil)
	transTypeUseCase := NewReadTransTypeUseCase(suite.transTypeRepoMock)
	transTypes, err := transTypeUseCase.ReadAllTransType()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyTransTypes, transTypes)
}

func TestTransTypePriceUseCaseSuite(t *testing.T) {
	suite.Run(t, new(TransTypeUseCaseTestSuite))
}
