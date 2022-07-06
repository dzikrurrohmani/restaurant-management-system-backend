package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/model"
	table_usecase "livecode-wmb-rest-api/usecase/table"
	"livecode-wmb-rest-api/utils"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type TableController struct {
	router       *gin.Engine
	ucTableCreate table_usecase.CreateTableUseCase
	ucTableRead   table_usecase.ReadTableUseCase
	ucTableUpdate table_usecase.UpdateTableUseCase
	ucTableDelete table_usecase.DeleteTableUseCase
	api.BaseApi
}

func (p *TableController) createNewTable(c *gin.Context) {
	var newTable model.Table
	err := p.ParseRequestBody(c, &newTable)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucTableCreate.CreateTable([]*model.Table{&newTable})
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newTable)
}

func (p *TableController) findAllTable(c *gin.Context) {
	Tables, err := p.ucTableRead.ReadAllTable()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, Tables)
}

func (p *TableController) updateTable(c *gin.Context) {
	var updatedTable model.Table
	err := p.ParseRequestBody(c, &updatedTable)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari table yg ingin diupdate
	tablewithID, err := p.ucTableRead.ReadTableById(updatedTable.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	// melakukan update
	updatedTableMap := structs.Map(updatedTable)
	err = p.ucTableUpdate.UpdateTable(&tablewithID, updatedTableMap)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, updatedTable)
}

func (p *TableController) deleteTable(c *gin.Context) {
	var deletedTable model.Table
	err := p.ParseRequestBody(c, &deletedTable)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari table yg ingin diupdate
	tablewithID, err := p.ucTableRead.ReadTableById(deletedTable.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	err = p.ucTableDelete.DeleteTable(&tablewithID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, deletedTable)
}

func NewTableController(
	router *gin.Engine,
	ucTableCreate table_usecase.CreateTableUseCase,
	ucTableRead table_usecase.ReadTableUseCase,
	ucTableUpdate table_usecase.UpdateTableUseCase,
	ucTableDelete table_usecase.DeleteTableUseCase) *TableController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := TableController{
		router:       router,
		ucTableCreate: ucTableCreate,
		ucTableRead:   ucTableRead,
		ucTableUpdate: ucTableUpdate,
		ucTableDelete: ucTableDelete,
	}

	// ini method-method nya
	router.POST("/table", controller.createNewTable)
	router.GET("/table", controller.findAllTable)
	router.PUT("/table", controller.updateTable)
	router.DELETE("/table", controller.deleteTable)
	return &controller
}
