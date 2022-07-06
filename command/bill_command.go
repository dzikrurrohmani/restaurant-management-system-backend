package command

import (
	"livecode-wmb-rest-api/utils"
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	usecase "livecode-wmb-rest-api/usecase/bill"
	"time"
)

func BillCreate(billRepo repository.BillRepository, tableRepo repository.TableRepository) {
	// Create (C)
	customerOrderUseCase := usecase.NewCustomerOrderUseCase(billRepo, tableRepo)
	billDetail01 := model.BillDetail{
		MenuPriceID: 1,
		Qty:         5,
	}
	billDetail02 := model.BillDetail{
		MenuPriceID: 2,
		Qty:         3,
	}
	bill01 := model.Bill{
		TransDate:   time.Now(),
		CustomerID:  1,
		TableID:     2,
		TransTypeID: "EI",
		BillDetails: []model.BillDetail{
			billDetail01,
			billDetail02,
		},
	}
	err := customerOrderUseCase.CreateOrder(&bill01)
	utils.RaiseError(err)
}

func BillPayment(billRepo repository.BillRepository, tableRepo repository.TableRepository) {
	// Update (U)
	customerPaymentUseCase := usecase.NewCustomerPaymentUseCase(billRepo, tableRepo)
	err := customerPaymentUseCase.OrderPayment(3)
	utils.RaiseError(err)
}
