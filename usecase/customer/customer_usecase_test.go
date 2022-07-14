package customer_usecase

import (
	"testing"

	"livecode-wmb-2/model"
	"livecode-wmb-2/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dummyCustomers = []model.Customer{
	{
		ID:            1,
		CustomerName:  "Dummy Name 1",
		MobilePhoneNo: "081xxx",
		IsMember:      false,
		Discounts:     nil,
	},
	{
		ID:            1,
		CustomerName:  "Dummy Name 1",
		MobilePhoneNo: "081xxx",
		IsMember:      true,
		Discounts:     nil,
	},
}

type discountRepoMock struct {
	mock.Mock
}

func (r *discountRepoMock) Create(discount []*model.Discount) error { panic("") }
func (r *discountRepoMock) FindBy(by map[string]interface{}) ([]model.Discount, error) {
	args := r.Called(by)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Discount), nil
}
func (r *discountRepoMock) FindAll() ([]model.Discount, error) { panic("") }
func (r *discountRepoMock) UpdateBy(discount *model.Discount, by map[string]interface{}) error {
	panic("")
}
func (r *discountRepoMock) Delete(discount *model.Discount) error { panic("") }

type customerRepoMock struct {
	mock.Mock
}

func (r *customerRepoMock) Create(customer []*model.Customer) error {
	args := r.Called(customer)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}
func (r *customerRepoMock) FindBy(by map[string]interface{}) ([]model.Customer, error) {
	args := r.Called(by)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]model.Customer), nil
}
func (r *customerRepoMock) UpdateBy(customer *model.Customer, by map[string]interface{}) error {
	args := r.Called(customer, by)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *customerRepoMock) UpdateAssociation(assocModel *model.Customer, assocName string, assocNewValue interface{}) error {
	args := r.Called(assocModel, assocName, assocNewValue)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (r *customerRepoMock) FindAll() ([]model.Customer, error)    { panic("") }
func (r *customerRepoMock) Delete(customer *model.Customer) error { panic("") }

type CustomerUseCaseTestSuite struct {
	suite.Suite
	customerRepoMock *customerRepoMock
	discountRepoMock *discountRepoMock
}

func (suite *CustomerUseCaseTestSuite) SetupTest() {
	suite.customerRepoMock = new(customerRepoMock)
	suite.discountRepoMock = new(discountRepoMock)
}

func (suite *CustomerUseCaseTestSuite) TestRegisterCustomer() {
	dummyCustomer := dummyCustomers[0]
	suite.customerRepoMock.On("Create", []*model.Customer{&dummyCustomer}).Return(nil) // expected, Get(0) : dummyCustomer, Get(1) : nil
	customerUseCase := NewCustomerRegistrationUseCase(suite.customerRepoMock)
	err := customerUseCase.CreateCustomer([]*model.Customer{&dummyCustomer}) // actual
	assert.Nil(suite.T(), err)
}

func (suite *CustomerUseCaseTestSuite) TestActivateMember_Success() {
	dummyCustomer := dummyCustomers[0]
	dummyCustomer2 := dummyCustomers[1]
	discount := model.Discount{
		ID:          1,
		Description: "mendapat potongan harga sebanyak 10 persen dari total belanja",
		Pct:         10,
	}
	by := map[string]interface{}{"mobile_phone_no": dummyCustomer.MobilePhoneNo}
	suite.customerRepoMock.On("FindBy", by).Return([]model.Customer{dummyCustomer}, nil).Once()                  // expected, Get(0) : dummyCustomer, Get(1) : nil
	suite.customerRepoMock.On("UpdateBy", &dummyCustomer, map[string]interface{}{"is_member": true}).Return(nil) // expected, Get(0) : dummyCustomer, Get(1) : nil
	suite.discountRepoMock.On("FindBy", map[string]interface{}{"id": discount.ID}).Return([]model.Discount{discount}, nil)
	suite.customerRepoMock.On("UpdateAssociation", &dummyCustomer, "Discounts", &discount).Return(nil) // expected, Get(0) : dummyCustomer, Get(1) : nil
	suite.customerRepoMock.On("FindBy", by).Return([]model.Customer{dummyCustomer2}, nil).Once()       // expected, Get(0) : dummyCustomer, Get(1) : nil
	customerUseCase := NewMemberActivationUseCase(suite.customerRepoMock, suite.discountRepoMock)
	customer, err := customerUseCase.ActivateMember(dummyCustomer.MobilePhoneNo, discount.ID) // actual
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), true, customer.IsMember)
}

func (suite *CustomerUseCaseTestSuite) TestActivateMember_Failed1() {
	dummyCustomer := dummyCustomers[0]
	by := map[string]interface{}{"mobile_phone_no": dummyCustomer.MobilePhoneNo}
	suite.customerRepoMock.On("FindBy", by).Return(nil, utils.DataNotFoundError()) // expected, Get(0) : dummyCustomer, Get(1) : nil
	customerUseCase := NewMemberActivationUseCase(suite.customerRepoMock, suite.discountRepoMock)
	customer, err := customerUseCase.ActivateMember(dummyCustomer.MobilePhoneNo, 1) // actual
	// assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), model.Customer{}, customer)
	assert.Equal(suite.T(), utils.DataNotFoundError(), err)
}

func (suite *CustomerUseCaseTestSuite) TestActivateMember_Failed2() {
	dummyCustomer := dummyCustomers[1]
	by := map[string]interface{}{"mobile_phone_no": dummyCustomer.MobilePhoneNo}
	suite.customerRepoMock.On("FindBy", by).Return([]model.Customer{dummyCustomer}, nil)                 // expected, Get(0) : dummyCustomer, Get(1) : nil
	customerUseCase := NewMemberActivationUseCase(suite.customerRepoMock, suite.discountRepoMock)
	customer, err := customerUseCase.ActivateMember(dummyCustomer.MobilePhoneNo, 1) // actual
	assert.Equal(suite.T(), dummyCustomer, customer)
	assert.NotNil(suite.T(), err)
}

func TestCustomerPriceUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CustomerUseCaseTestSuite))
}
