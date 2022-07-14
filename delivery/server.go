package delivery

import (
	"livecode-wmb-2/config"
	"livecode-wmb-2/delivery/controller"
	"livecode-wmb-2/manager"

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
	host := appConfig.ApiUrl
	return &appServer{
		useCaseManager: useCaseManager,
		engine:         r,
		host:           host,
	}
}

func (a *appServer) initControllers() {
	controller.NewMenuController(a.engine, a.useCaseManager.CreateMenuUseCase(), a.useCaseManager.ReadMenuUseCase(), a.useCaseManager.UpdateMenuUseCase(), a.useCaseManager.DeleteMenuUseCase())
	controller.NewMenuPriceController(a.engine, a.useCaseManager.CreateMenuPriceUseCase(), a.useCaseManager.ReadMenuPriceUseCase(), a.useCaseManager.UpdateMenuPriceUseCase(), a.useCaseManager.DeleteMenuPriceUseCase())
	controller.NewDiscountController(a.engine, a.useCaseManager.CreateDiscountUseCase(), a.useCaseManager.ReadDiscountUseCase(), a.useCaseManager.UpdateDiscountUseCase(), a.useCaseManager.DeleteDiscountUseCase())
	controller.NewTableController(a.engine, a.useCaseManager.CreateTableUseCase(), a.useCaseManager.ReadTableUseCase(), a.useCaseManager.UpdateTableUseCase(), a.useCaseManager.DeleteTableUseCase())
	controller.NewTransTypeController(a.engine, a.useCaseManager.CreateTransTypeUseCase(), a.useCaseManager.ReadTransTypeUseCase(), a.useCaseManager.UpdateTransTypeUseCase(), a.useCaseManager.DeleteTransTypeUseCase())
	controller.NewCustomerController(a.engine, a.useCaseManager.CustomerRegistrationUseCase(), a.useCaseManager.MemberActivationUseCase())
	controller.NewBillController(a.engine, a.useCaseManager.CustomerOrderUseCase(), a.useCaseManager.CustomerPaymentUseCase())
	controller.NewBillDetailController(a.engine, a.useCaseManager.GetIncomeUseCase())
}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
