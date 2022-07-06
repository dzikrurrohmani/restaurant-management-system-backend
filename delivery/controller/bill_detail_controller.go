package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	billdetail_usecase "livecode-wmb-rest-api/usecase/bill_detail"
	"time"

	"github.com/gin-gonic/gin"
)

type BillDetailController struct {
	router      *gin.Engine
	ucGetIncome billdetail_usecase.GetIncomeUseCase
	api.BaseApi
}

func (p *BillDetailController) getIncome(c *gin.Context) {
	var result []struct {
		Date  time.Time
		Total int64
	}
	err := p.ucGetIncome.GetDailyIncome(&result)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, result)
}

func NewBillDetailController(
	router *gin.Engine,
	ucGetIncome billdetail_usecase.GetIncomeUseCase) *BillDetailController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := BillDetailController{
		router:      router,
		ucGetIncome: ucGetIncome,
	}

	// ini method-method nya
	router.GET("/billdetail", controller.getIncome)
	return &controller
}
