package controller

import (
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/model"
	menuprice_usecase "livecode-wmb-2/usecase/menu_price"
	"livecode-wmb-2/utils"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type MenuPriceController struct {
	router            *gin.Engine
	ucMenuPriceCreate menuprice_usecase.CreateMenuPriceUseCase
	ucMenuPriceRead   menuprice_usecase.ReadMenuPriceUseCase
	ucMenuPriceUpdate menuprice_usecase.UpdateMenuPriceUseCase
	ucMenuPriceDelete menuprice_usecase.DeleteMenuPriceUseCase
	api.BaseApi
}

func (p *MenuPriceController) createNewMenuPrice(c *gin.Context) {
	var newMenuPrice model.MenuPrice
	err := p.ParseRequestBody(c, &newMenuPrice)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucMenuPriceCreate.CreateMenuPrice([]*model.MenuPrice{&newMenuPrice})
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newMenuPrice)
}

func (p *MenuPriceController) findAllMenuPrice(c *gin.Context) {
	MenuPrices, err := p.ucMenuPriceRead.ReadAllMenuPrice()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, MenuPrices)
}

func (p *MenuPriceController) updateMenuPrice(c *gin.Context) {
	var updatedMenuPrice model.MenuPrice
	err := p.ParseRequestBody(c, &updatedMenuPrice)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari menuprice yg ingin diupdate
	menupricewithID, err := p.ucMenuPriceRead.ReadMenuPriceById(updatedMenuPrice.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	// melakukan update
	updatedMenuPriceMap := structs.Map(updatedMenuPrice)
	err = p.ucMenuPriceUpdate.UpdateMenuPrice(&menupricewithID, updatedMenuPriceMap)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, updatedMenuPrice)
}

func (p *MenuPriceController) deleteMenuPrice(c *gin.Context) {
	var deletedMenuPrice model.MenuPrice
	err := p.ParseRequestBody(c, &deletedMenuPrice)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	// cari menuprice yg ingin diupdate
	menupricewithID, err := p.ucMenuPriceRead.ReadMenuPriceById(deletedMenuPrice.ID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	err = p.ucMenuPriceDelete.DeleteMenuPrice(&menupricewithID)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, deletedMenuPrice)
}

func NewMenuPriceController(
	router *gin.Engine,
	ucMenuPriceCreate menuprice_usecase.CreateMenuPriceUseCase,
	ucMenuPriceRead menuprice_usecase.ReadMenuPriceUseCase,
	ucMenuPriceUpdate menuprice_usecase.UpdateMenuPriceUseCase,
	ucMenuPriceDelete menuprice_usecase.DeleteMenuPriceUseCase) *MenuPriceController {
	// Disini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := MenuPriceController{
		router:            router,
		ucMenuPriceCreate: ucMenuPriceCreate,
		ucMenuPriceRead:   ucMenuPriceRead,
		ucMenuPriceUpdate: ucMenuPriceUpdate,
		ucMenuPriceDelete: ucMenuPriceDelete,
	}

	// ini method-method nya
	router.POST("/menuprice", controller.createNewMenuPrice)
	router.GET("/menuprice", controller.findAllMenuPrice)
	router.PUT("/menuprice", controller.updateMenuPrice)
	router.DELETE("/menuprice", controller.deleteMenuPrice)
	return &controller
}
