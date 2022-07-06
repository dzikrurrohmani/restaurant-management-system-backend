package delivery

import (
	"livecode-wmb-rest-api/config"
	"livecode-wmb-rest-api/delivery/controller"
	"livecode-wmb-rest-api/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	r := gin.Default()
	appConfig := config.NewConfig()
	infra := manager.NewInfra(appConfig)
	repoManager := manager.NewRepositoryManager(infra) // kalo misal slice (di sini isinya slice)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := appConfig.Url
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}

func (a *appServer) initControllers() {
	// controller.NewProductController(a.engine, a.useCaseManager.Cre, a.useCaseManager.ListProductUseCase())
	controller.NewMenuController(a.engine, a.useCaseManager.CreateMenuUseCase(), a.useCaseManager.ReadMenuUseCase(), a.useCaseManager.UpdateMenuUseCase(), a.useCaseManager.DeleteMenuUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
