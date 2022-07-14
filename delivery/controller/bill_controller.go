package controller

import (
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/model"
	bill_usecase "livecode-wmb-2/usecase/bill"
	"livecode-wmb-2/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type BillController struct {
	router        *gin.Engine
	ucCustOrder   bill_usecase.CustomerOrderUseCase
	ucCustPayment bill_usecase.CustomerPaymentUseCase
	api.BaseApi
}

func (p *BillController) custOrder(c *gin.Context) {
	var newBill model.Bill
	err := p.ParseRequestBody(c, &newBill)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	newBill.TransDate = time.Now()
	err = p.ucCustOrder.CreateOrder(&newBill)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newBill)
}

func (p *BillController) custPayment(c *gin.Context) {
	type custPaymentInput struct {
		BillId uint
	}
	var temp custPaymentInput
	err := p.ParseRequestBody(c, &temp)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	Bill, err := p.ucCustPayment.OrderPayment(temp.BillId)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, Bill)
}

func NewBillController(
	router *gin.Engine,
	ucCustOrder bill_usecase.CustomerOrderUseCase,
	ucCustPayment bill_usecase.CustomerPaymentUseCase) *BillController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := BillController{
		router:        router,
		ucCustOrder:   ucCustOrder,
		ucCustPayment: ucCustPayment,
	}

	// ini method-method nya
	router.POST("/bill", controller.custOrder)
	router.GET("/bill", controller.custPayment)
	return &controller
}
