package controller

import (
	"livecode-wmb-rest-api/delivery/api"
	"livecode-wmb-rest-api/model"
	menu_usecase "livecode-wmb-rest-api/usecase/menu"
	"livecode-wmb-rest-api/utils"

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
	menuwithID, err := p.ucMenuRead.ReadMenuById(updatedMenu.BaseModel.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	// melakukan update
	updatedMenuMap := structs.Map(updatedMenu)
	p.ucMenuUpdate.UpdateMenu(&menuwithID, map[string]interface{})

	err = p.ucMenu.CreateMenu(&updatedMenu)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, updatedMenu)
	// // Misal ganti nama
	// UpdateMenuUseCase := usecase.NewUpdateMenuUseCase(menuRepo)
	// UpdateMenuUseCase.UpdateMenu(&menuwithID, map[string]interface{}{"menu_name": "Nasi Liwet"})

	// // Dicetak kembali untuk dicek apakah sudah berubah
	// menuwithID, err = ReadMenuUseCase.ReadMenuById(3)
	// utils.RaiseError(err)
	// log.Println(menuwithID)
}

func (p *MenuController) deleteMenu(c *gin.Context) {
	var newMenu model.Menu
	err := p.ParseRequestBody(c, &newMenu)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucMenu.CreateMenu(&newMenu)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newMenu)
}

func NewMenuController(
	router *gin.Engine,
	ucMenu usecase.CreateMenuUseCase,
	ucMenuList usecase.ListMenuUseCase) *MenuController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := MenuController{
		router:     router,
		ucMenu:     ucMenu,
		ucMenuList: ucMenuList,
	}

	// ini method-method nya
	router.POST("/Menu", controller.createNewMenu)
	router.GET("/Menu", controller.findAllMenu)
	return &controller
}
