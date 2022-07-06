package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/model"
	transtype_usecase "livecode-wmb-rest-api/usecase/trans_type"
	"livecode-wmb-rest-api/utils"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type TransTypeController struct {
	router       *gin.Engine
	ucTransTypeCreate transtype_usecase.CreateTransTypeUseCase
	ucTransTypeRead   transtype_usecase.ReadTransTypeUseCase
	ucTransTypeUpdate transtype_usecase.UpdateTransTypeUseCase
	ucTransTypeDelete transtype_usecase.DeleteTransTypeUseCase
	api.BaseApi
}

func (p *TransTypeController) createNewTransType(c *gin.Context) {
	var newTransType model.TransType
	err := p.ParseRequestBody(c, &newTransType)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucTransTypeCreate.CreateTransType([]*model.TransType{&newTransType})
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newTransType)
}

func (p *TransTypeController) findAllTransType(c *gin.Context) {
	TransTypes, err := p.ucTransTypeRead.ReadAllTransType()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, TransTypes)
}

func (p *TransTypeController) updateTransType(c *gin.Context) {
	var updatedTransType model.TransType
	err := p.ParseRequestBody(c, &updatedTransType)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari transtype yg ingin diupdate
	transtypewithID, err := p.ucTransTypeRead.ReadTransTypeById(updatedTransType.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	// melakukan update
	updatedTransTypeMap := structs.Map(updatedTransType)
	err = p.ucTransTypeUpdate.UpdateTransType(&transtypewithID, updatedTransTypeMap)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, updatedTransType)
}

func (p *TransTypeController) deleteTransType(c *gin.Context) {
	var deletedTransType model.TransType
	err := p.ParseRequestBody(c, &deletedTransType)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari transtype yg ingin diupdate
	transtypewithID, err := p.ucTransTypeRead.ReadTransTypeById(deletedTransType.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	err = p.ucTransTypeDelete.DeleteTransType(&transtypewithID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, deletedTransType)
}

func NewTransTypeController(
	router *gin.Engine,
	ucTransTypeCreate transtype_usecase.CreateTransTypeUseCase,
	ucTransTypeRead transtype_usecase.ReadTransTypeUseCase,
	ucTransTypeUpdate transtype_usecase.UpdateTransTypeUseCase,
	ucTransTypeDelete transtype_usecase.DeleteTransTypeUseCase) *TransTypeController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := TransTypeController{
		router:       router,
		ucTransTypeCreate: ucTransTypeCreate,
		ucTransTypeRead:   ucTransTypeRead,
		ucTransTypeUpdate: ucTransTypeUpdate,
		ucTransTypeDelete: ucTransTypeDelete,
	}

	// ini method-method nya
	router.POST("/transtype", controller.createNewTransType)
	router.GET("/transtype", controller.findAllTransType)
	router.PUT("/transtype", controller.updateTransType)
	router.DELETE("/transtype", controller.deleteTransType)
	return &controller
}
