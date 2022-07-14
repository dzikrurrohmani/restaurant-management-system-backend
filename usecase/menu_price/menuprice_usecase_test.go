package menuprice_usecase

import (
	"errors"
	"testing"

	"livecode-wmb-2/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMenuPrices = []model.MenuPrice{
	{
		ID:     1,
		Price:  20000,
		MenuID: nil,
	},
	{
		ID:     2,
		Price:  20000,
		MenuID: nil,
	},
}

type menupriceRepoMock struct {
	mock.Mock
}

type MenuPriceUseCaseTestSuite struct {
	suite.Suite
	menupriceRepoMock *menupriceRepoMock
	// menupriceRepoMock repository.MenuPriceRepository
}

func (r *menupriceRepoMock) Create(menuprice []*model.MenuPrice) error {
	args := r.Called(menuprice)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *menupriceRepoMock) FindBy(by map[string]interface{}) ([]model.MenuPrice, error) {
	args := r.Called(by)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.MenuPrice), nil
}
func (r *menupriceRepoMock) FindAll() ([]model.MenuPrice, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.MenuPrice), nil
}
func (r *menupriceRepoMock) UpdateBy(menuprice *model.MenuPrice, by map[string]interface{}) error {
	args := r.Called(menuprice, by)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *menupriceRepoMock) Delete(menuprice *model.MenuPrice) error {
	args := r.Called(menuprice)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite *MenuPriceUseCaseTestSuite) SetupTest() {
	suite.menupriceRepoMock = new(menupriceRepoMock)
}

func (suite *MenuPriceUseCaseTestSuite) TestCreateMenuPrice() {
	dummyMenuPrice := dummyMenuPrices[0]
	suite.menupriceRepoMock.On("Create", []*model.MenuPrice{&dummyMenuPrice}).Return(nil) // expected, Get(0) : dummyMenuPrice, Get(1) : nil
	menupriceUseCase := NewCreateMenuPriceUseCase(suite.menupriceRepoMock)
	err := menupriceUseCase.CreateMenuPrice([]*model.MenuPrice{&dummyMenuPrice}) // actual
	assert.Nil(suite.T(), err)
}

func (suite *MenuPriceUseCaseTestSuite) TestDeleteMenuPrice() {
	dummyMenuPrice := dummyMenuPrices[0]
	suite.menupriceRepoMock.On("Delete", &dummyMenuPrice).Return(nil) // expected, Get(0) : dummyMenuPrice, Get(1) : nil
	menupriceUseCase := NewDeleteMenuPriceUseCase(suite.menupriceRepoMock)
	err := menupriceUseCase.DeleteMenuPrice(&dummyMenuPrice) // actual
	assert.Nil(suite.T(), err)
}

func (suite *MenuPriceUseCaseTestSuite) TestUpdateMenuPrice() {
	var by = map[string]interface{}{
		"Id": 1,
	}
	dummyMenuPrice := dummyMenuPrices[0]
	suite.menupriceRepoMock.On("UpdateBy", &dummyMenuPrice, by).Return(nil) // expected, Get(0) : dummyMenuPrice, Get(1) : nil
	menupriceUseCase := NewUpdateMenuPriceUseCase(suite.menupriceRepoMock)
	err := menupriceUseCase.UpdateMenuPrice(&dummyMenuPrice, by) // actual
	assert.Nil(suite.T(), err)
}

func (suite *MenuPriceUseCaseTestSuite) TestMenuPriceFindById_Success() {
	dummyMenuPrice := dummyMenuPrices[0]
	// Mengacu ke nama method
	suite.menupriceRepoMock.On("FindBy", map[string]interface{}{"id": dummyMenuPrice.ID}).Return([]model.MenuPrice{dummyMenuPrice}, nil)
	menupriceUseCase := NewReadMenuPriceUseCase(suite.menupriceRepoMock)
	menuprice, err := menupriceUseCase.ReadMenuPriceById(dummyMenuPrice.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenuPrice.ID, menuprice.ID)
}

func (suite *MenuPriceUseCaseTestSuite) TestMenuPriceFindById_Failed() {
	dummyMenuPrice := dummyMenuPrices[0]
	// Mengacu ke nama method
	suite.menupriceRepoMock.On("FindBy", map[string]interface{}{"id": dummyMenuPrice.ID}).Return(nil, errors.New("failed"))
	menupriceUseCase := NewReadMenuPriceUseCase(suite.menupriceRepoMock)
	menuprice, err := menupriceUseCase.ReadMenuPriceById(dummyMenuPrice.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Equal(suite.T(), model.MenuPrice{}, menuprice)
}

func (suite *MenuPriceUseCaseTestSuite) TestMenuPriceFindBy() {
	dummyMenuPrice := dummyMenuPrices[0]
	var by = map[string]interface{}{"id": dummyMenuPrice.ID}
	// Mengacu ke nama method
	suite.menupriceRepoMock.On("FindBy", by).Return([]model.MenuPrice{dummyMenuPrice}, nil)
	menupriceUseCase := NewReadMenuPriceUseCase(suite.menupriceRepoMock)
	menuprice, err := menupriceUseCase.ReadMenuPriceBy(by)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenuPrice.ID, menuprice[0].ID)
}

func (suite *MenuPriceUseCaseTestSuite) TestMenuPriceReadeAll() {
	// Mengacu ke nama method
	suite.menupriceRepoMock.On("FindAll").Return(dummyMenuPrices, nil)
	menupriceUseCase := NewReadMenuPriceUseCase(suite.menupriceRepoMock)
	menus, err := menupriceUseCase.ReadAllMenuPrice()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenuPrices, menus)
}

func TestMenuPricePriceUseCaseSuite(t *testing.T) {
	suite.Run(t, new(MenuPriceUseCaseTestSuite))
}
