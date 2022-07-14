package menu_usecase

import (
	"errors"
	"testing"

	"livecode-wmb-2/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyMenus = []model.Menu{
	{
		ID:         1,
		MenuName:   "Dummy Name 1",
		MenuPrices: nil,
	},
	{
		ID:         2,
		MenuName:   "Dummy Name 2",
		MenuPrices: nil,
	},
}

type menuRepoMock struct {
	mock.Mock
}

type MenuUseCaseTestSuite struct {
	suite.Suite
	menuRepoMock *menuRepoMock
	// menuRepoMock repository.MenuRepository
}

func (r *menuRepoMock) Create(menu []*model.Menu) error {
	args := r.Called(menu)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *menuRepoMock) FindBy(by map[string]interface{}) ([]model.Menu, error) {
	args := r.Called(by)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Menu), nil
}
func (r *menuRepoMock) FindAll() ([]model.Menu, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Menu), nil
}
func (r *menuRepoMock) UpdateBy(menu *model.Menu, by map[string]interface{}) error {
	args := r.Called(menu, by)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *menuRepoMock) Delete(menu *model.Menu) error {
	args := r.Called(menu)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite *MenuUseCaseTestSuite) SetupTest() {
	suite.menuRepoMock = new(menuRepoMock)
}

func (suite *MenuUseCaseTestSuite) TestCreateMenu() {
	dummyMenu := dummyMenus[0]
	suite.menuRepoMock.On("Create", []*model.Menu{&dummyMenu}).Return(nil) // expected, Get(0) : dummyMenu, Get(1) : nil
	menuUseCaseTest := NewCreateMenuUseCase(suite.menuRepoMock)
	err := menuUseCaseTest.CreateMenu([]*model.Menu{&dummyMenu}) // actual
	assert.Nil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestDeleteMenu() {
	dummyMenu := dummyMenus[0]
	suite.menuRepoMock.On("Delete", &dummyMenu).Return(nil) // expected, Get(0) : dummyMenu, Get(1) : nil
	menuUseCaseTest := NewDeleteMenuUseCase(suite.menuRepoMock)
	err := menuUseCaseTest.DeleteMenu(&dummyMenu) // actual
	assert.Nil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestUpdateMenu() {
	var by = map[string]interface{}{
		"Id": 1,
	}
	dummyMenu := dummyMenus[0]
	suite.menuRepoMock.On("UpdateBy", &dummyMenu, by).Return(nil) // expected, Get(0) : dummyMenu, Get(1) : nil
	menuUseCaseTest := NewUpdateMenuUseCase(suite.menuRepoMock)
	err := menuUseCaseTest.UpdateMenu(&dummyMenu, by) // actual
	assert.Nil(suite.T(), err)
}

func (suite *MenuUseCaseTestSuite) TestMenuFindById_Success() {
	dummyMenu := dummyMenus[0]
	// Mengacu ke nama method
	suite.menuRepoMock.On("FindBy", map[string]interface{}{"id": dummyMenu.ID}).Return([]model.Menu{dummyMenu}, nil)
	menuUseCase := NewReadMenuUseCase(suite.menuRepoMock)
	menu, err := menuUseCase.ReadMenuById(dummyMenu.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenu.ID, menu.ID)
}

func (suite *MenuUseCaseTestSuite) TestMenuFindById_Failed() {
	dummyMenu := dummyMenus[0]
	// Mengacu ke nama method
	suite.menuRepoMock.On("FindBy", map[string]interface{}{"id": dummyMenu.ID}).Return(nil, errors.New("failed"))
	menuUseCase := NewReadMenuUseCase(suite.menuRepoMock)
	menu, err := menuUseCase.ReadMenuById(dummyMenu.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Equal(suite.T(), model.Menu{}, menu)
}

func (suite *MenuUseCaseTestSuite) TestMenuFindBy() {
	dummyMenu := dummyMenus[0]
	var by = map[string]interface{}{"id": dummyMenu.ID}
	// Mengacu ke nama method
	suite.menuRepoMock.On("FindBy", by).Return([]model.Menu{dummyMenu}, nil)
	menuUseCase := NewReadMenuUseCase(suite.menuRepoMock)
	menu, err := menuUseCase.ReadMenuBy(by)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenu.ID, menu[0].ID)
}

func (suite *MenuUseCaseTestSuite) TestMenuReadeAll() {
	// Mengacu ke nama method
	suite.menuRepoMock.On("FindAll").Return(dummyMenus, nil)
	menuUseCase := NewReadMenuUseCase(suite.menuRepoMock)
	menus, err := menuUseCase.ReadAllMenu()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyMenus, menus)
}

func TestMenuUseCaseSuite(t *testing.T) {
	suite.Run(t, new(MenuUseCaseTestSuite))
}
