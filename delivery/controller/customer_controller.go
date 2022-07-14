package controller

import (
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/model"
	customer_usecase "livecode-wmb-2/usecase/customer"
	"livecode-wmb-2/utils"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router                 *gin.Engine
	ucCustomerRegistration customer_usecase.CustomerRegistrationUseCase
	ucCustomerActivation   customer_usecase.MemberActivationUseCase
	api.BaseApi
}

func (p *CustomerController) custRegist(c *gin.Context) {
	var newCustomer model.Customer
	err := p.ParseRequestBody(c, &newCustomer)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucCustomerRegistration.CreateCustomer([]*model.Customer{&newCustomer})
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newCustomer)
}

func (p *CustomerController) memberActivation(c *gin.Context) {
	type memberActivationInput struct {
		MobilePhoneNo string
		DiscountId    uint
	}
	var temp memberActivationInput
	err := p.ParseRequestBody(c, &temp)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	Customer, err := p.ucCustomerActivation.ActivateMember(temp.MobilePhoneNo, temp.DiscountId)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, Customer)
}

func NewCustomerController(
	router *gin.Engine,
	ucCustomerRegistration customer_usecase.CustomerRegistrationUseCase,
	ucCustomerActivation customer_usecase.MemberActivationUseCase) *CustomerController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := CustomerController{
		router:                 router,
		ucCustomerRegistration: ucCustomerRegistration,
		ucCustomerActivation:   ucCustomerActivation,
	}

	// ini method-method nya
	router.POST("/customer", controller.custRegist)
	router.PUT("/customer", controller.memberActivation)
	return &controller
}
