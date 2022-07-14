package controller

import (
	"livecode-wmb-2/delivery/api"
	"livecode-wmb-2/delivery/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	router *gin.Engine
	api.BaseApi
}

func (p *LoginController) appLogin(c *gin.Context) {
	var user middleware.Credential
	//authHeader := AuthHeader{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	if user.Username == "admin" && user.Password == "admin" {
		token, err := middleware.GenerateToken(user.Username, "admin@wmb.com")
		if err != nil {
			c.AbortWithStatus(401)
			return

		}
		c.JSON(200, gin.H{
			"token": token,
		})
	} else {
		c.AbortWithStatus(401)
	}
}

func NewLoginController(router *gin.Engine) *LoginController {
	controller := LoginController{
		router: router,
	}

	// ini method-method nya
	router.POST("/login", controller.appLogin)
	return &controller
}
