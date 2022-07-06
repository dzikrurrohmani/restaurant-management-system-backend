package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/model"
	menu_usecase "livecode-wmb-rest-api/usecase/menu"
	"livecode-wmb-rest-api/utils"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type MenuController struct {
	router       *gin.Engine
	ucMenuCreate menu_usecase.CreateMenuUseCase
	ucMenuRead   menu_usecase.ReadMenuUseCase
	ucMenuUpdate menu_usecase.UpdateMenuUseCase
	ucMenuDelete menu_usecase.DeleteMenuUseCase
	api.BaseApi
}

func (p *MenuController) createNewMenu(c *gin.Context) {
	var newMenu model.Menu
	err := p.ParseRequestBody(c, &newMenu)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucMenuCreate.CreateMenu([]*model.Menu{&newMenu})
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newMenu)
}

func (p *MenuController) findAllMenu(c *gin.Context) {
	Menus, err := p.ucMenuRead.ReadAllMenu()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, Menus)
}

func (p *MenuController) updateMenu(c *gin.Context) {
	var updatedMenu model.Menu
	err := p.ParseRequestBody(c, &updatedMenu)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari menu yg ingin diupdate
	menuwithID, err := p.ucMenuRead.ReadMenuById(updatedMenu.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	// melakukan update
	updatedMenuMap := structs.Map(updatedMenu)
	err = p.ucMenuUpdate.UpdateMenu(&menuwithID, updatedMenuMap)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, updatedMenu)
}

func (p *MenuController) deleteMenu(c *gin.Context) {
	var deletedMenu model.Menu
	err := p.ParseRequestBody(c, &deletedMenu)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari menu yg ingin diupdate
	menuwithID, err := p.ucMenuRead.ReadMenuById(deletedMenu.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	err = p.ucMenuDelete.DeleteMenu(&menuwithID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, deletedMenu)
}

func NewMenuController(
	router *gin.Engine,
	ucMenuCreate menu_usecase.CreateMenuUseCase,
	ucMenuRead menu_usecase.ReadMenuUseCase,
	ucMenuUpdate menu_usecase.UpdateMenuUseCase,
	ucMenuDelete menu_usecase.DeleteMenuUseCase) *MenuController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := MenuController{
		router:       router,
		ucMenuCreate: ucMenuCreate,
		ucMenuRead:   ucMenuRead,
		ucMenuUpdate: ucMenuUpdate,
		ucMenuDelete: ucMenuDelete,
	}

	// ini method-method nya
	router.POST("/menu", controller.createNewMenu)
	router.GET("/menu", controller.findAllMenu)
	router.PUT("/menu", controller.updateMenu)
	router.DELETE("/menu", controller.deleteMenu)
	return &controller
}
