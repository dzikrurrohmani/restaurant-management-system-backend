package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/model"
	discount_usecase "livecode-wmb-rest-api/usecase/discount"
	"livecode-wmb-rest-api/utils"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type DiscountController struct {
	router       *gin.Engine
	ucDiscountCreate discount_usecase.CreateDiscountUseCase
	ucDiscountRead   discount_usecase.ReadDiscountUseCase
	ucDiscountUpdate discount_usecase.UpdateDiscountUseCase
	ucDiscountDelete discount_usecase.DeleteDiscountUseCase
	api.BaseApi
}

func (p *DiscountController) createNewDiscount(c *gin.Context) {
	var newDiscount model.Discount
	err := p.ParseRequestBody(c, &newDiscount)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucDiscountCreate.CreateDiscount([]*model.Discount{&newDiscount})
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newDiscount)
}

func (p *DiscountController) findAllDiscount(c *gin.Context) {
	Discounts, err := p.ucDiscountRead.ReadAllDiscount()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, Discounts)
}

func (p *DiscountController) updateDiscount(c *gin.Context) {
	var updatedDiscount model.Discount
	err := p.ParseRequestBody(c, &updatedDiscount)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari discount yg ingin diupdate
	discountwithID, err := p.ucDiscountRead.ReadDiscountById(updatedDiscount.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	// melakukan update
	updatedDiscountMap := structs.Map(updatedDiscount)
	err = p.ucDiscountUpdate.UpdateDiscount(&discountwithID, updatedDiscountMap)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, updatedDiscount)
}

func (p *DiscountController) deleteDiscount(c *gin.Context) {
	var deletedDiscount model.Discount
	err := p.ParseRequestBody(c, &deletedDiscount)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari discount yg ingin diupdate
	discountwithID, err := p.ucDiscountRead.ReadDiscountById(deletedDiscount.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	err = p.ucDiscountDelete.DeleteDiscount(&discountwithID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, deletedDiscount)
}

func NewDiscountController(
	router *gin.Engine,
	ucDiscountCreate discount_usecase.CreateDiscountUseCase,
	ucDiscountRead discount_usecase.ReadDiscountUseCase,
	ucDiscountUpdate discount_usecase.UpdateDiscountUseCase,
	ucDiscountDelete discount_usecase.DeleteDiscountUseCase) *DiscountController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := DiscountController{
		router:       router,
		ucDiscountCreate: ucDiscountCreate,
		ucDiscountRead:   ucDiscountRead,
		ucDiscountUpdate: ucDiscountUpdate,
		ucDiscountDelete: ucDiscountDelete,
	}

	// ini method-method nya
	router.POST("/discount", controller.createNewDiscount)
	router.GET("/discount", controller.findAllDiscount)
	router.PUT("/discount", controller.updateDiscount)
	router.DELETE("/discount", controller.deleteDiscount)
	return &controller
}
