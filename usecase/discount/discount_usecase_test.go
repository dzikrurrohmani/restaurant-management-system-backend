package discount_usecase

import (
	"errors"
	"testing"

	"livecode-wmb-2/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyDiscounts = []model.Discount{
	{
		ID:          1,
		Description: "mendapat potongan harga sebanyak 10 persen dari total belanja",
		Pct:         10,
	},
	{
		ID:          2,
		Description: "mendapat potongan harga sebanyak 5 persen dari total belanja",
		Pct:         5,
	},
}

type discountRepoMock struct {
	mock.Mock
}

type DiscountUseCaseTestSuite struct {
	suite.Suite
	discountRepoMock *discountRepoMock
	// discountRepoMock repository.DiscountRepository
}

func (r *discountRepoMock) Create(discount []*model.Discount) error {
	args := r.Called(discount)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *discountRepoMock) FindBy(by map[string]interface{}) ([]model.Discount, error) {
	args := r.Called(by)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Discount), nil
}
func (r *discountRepoMock) FindAll() ([]model.Discount, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Discount), nil
}
func (r *discountRepoMock) UpdateBy(discount *model.Discount, by map[string]interface{}) error {
	args := r.Called(discount, by)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *discountRepoMock) Delete(discount *model.Discount) error {
	args := r.Called(discount)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (suite *DiscountUseCaseTestSuite) SetupTest() {
	suite.discountRepoMock = new(discountRepoMock)
}

func (suite *DiscountUseCaseTestSuite) TestCreateDiscount() {
	dummyDiscount := dummyDiscounts[0]
	suite.discountRepoMock.On("Create", []*model.Discount{&dummyDiscount}).Return(nil) // expected, Get(0) : dummyDiscount, Get(1) : nil
	discountUseCase := NewCreateDiscountUseCase(suite.discountRepoMock)
	err := discountUseCase.CreateDiscount([]*model.Discount{&dummyDiscount}) // actual
	assert.Nil(suite.T(), err)
}

func (suite *DiscountUseCaseTestSuite) TestDeleteDiscount() {
	dummyDiscount := dummyDiscounts[0]
	suite.discountRepoMock.On("Delete", &dummyDiscount).Return(nil) // expected, Get(0) : dummyDiscount, Get(1) : nil
	discountUseCase := NewDeleteDiscountUseCase(suite.discountRepoMock)
	err := discountUseCase.DeleteDiscount(&dummyDiscount) // actual
	assert.Nil(suite.T(), err)
}

func (suite *DiscountUseCaseTestSuite) TestUpdateDiscount() {
	var by = map[string]interface{}{
		"Id": 1,
	}
	dummyDiscount := dummyDiscounts[0]
	suite.discountRepoMock.On("UpdateBy", &dummyDiscount, by).Return(nil) // expected, Get(0) : dummyDiscount, Get(1) : nil
	discountUseCase := NewUpdateDiscountUseCase(suite.discountRepoMock)
	err := discountUseCase.UpdateDiscount(&dummyDiscount, by) // actual
	assert.Nil(suite.T(), err)
}

func (suite *DiscountUseCaseTestSuite) TestDiscountFindById_Success() {
	dummyDiscount := dummyDiscounts[0]
	// Mengacu ke nama method
	suite.discountRepoMock.On("FindBy", map[string]interface{}{"id": dummyDiscount.ID}).Return([]model.Discount{dummyDiscount}, nil)
	discountUseCase := NewReadDiscountUseCase(suite.discountRepoMock)
	discount, err := discountUseCase.ReadDiscountById(dummyDiscount.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyDiscount.ID, discount.ID)
}

func (suite *DiscountUseCaseTestSuite) TestDiscountFindById_Failed() {
	dummyDiscount := dummyDiscounts[0]
	// Mengacu ke nama method
	suite.discountRepoMock.On("FindBy", map[string]interface{}{"id": dummyDiscount.ID}).Return(nil, errors.New("failed"))
	discountUseCase := NewReadDiscountUseCase(suite.discountRepoMock)
	discount, err := discountUseCase.ReadDiscountById(dummyDiscount.ID)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "failed", err.Error())
	assert.Equal(suite.T(), model.Discount{}, discount)
}

func (suite *DiscountUseCaseTestSuite) TestDiscountFindBy() {
	dummyDiscount := dummyDiscounts[0]
	var by = map[string]interface{}{"id": dummyDiscount.ID}
	// Mengacu ke nama method
	suite.discountRepoMock.On("FindBy", by).Return([]model.Discount{dummyDiscount}, nil)
	discountUseCase := NewReadDiscountUseCase(suite.discountRepoMock)
	discount, err := discountUseCase.ReadDiscountBy(by)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyDiscount.ID, discount[0].ID)
}

func (suite *DiscountUseCaseTestSuite) TestDiscountReadeAll() {
	// Mengacu ke nama method
	suite.discountRepoMock.On("FindAll").Return(dummyDiscounts, nil)
	discountUseCase := NewReadDiscountUseCase(suite.discountRepoMock)
	discounts, err := discountUseCase.ReadAllDiscount()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), dummyDiscounts, discounts)
}

func TestDiscountPriceUseCaseSuite(t *testing.T) {
	suite.Run(t, new(DiscountUseCaseTestSuite))
}
