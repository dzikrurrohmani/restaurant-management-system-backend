package bill_usecase

import (
	"errors"
	"fmt"
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type CustomerOrderUseCase interface {
	CreateOrder(bill *model.Bill) error
}

type customerOrderUseCase struct {
	billRepo  repository.BillRepository
	tableRepo repository.TableRepository
}

func (c *customerOrderUseCase) CreateOrder(bill *model.Bill) error {
	if bill.TransTypeID == "EI" {
		tableSlice, err := c.tableRepo.FindBy(map[string]interface{}{"id": bill.TableID})
		if err != nil {
			fmt.Println("Informasi meja tidak ditemukan.")
			return err
		}
		tableSelected := tableSlice[0]
		if !tableSelected.IsAvailable {
			fmt.Println("Meja tidak tersedia")
			return errors.New("meja tidak tersedia")
		}
		tableSelected.IsAvailable = false
		err = c.billRepo.Create(bill, &tableSelected)
		if err != nil {
			fmt.Println("Bill gagal dibuat.")
			return err
		}
	} else {
		err := c.billRepo.Create(bill, nil)
		if err != nil {
			fmt.Println("Bill gagal dibuat.")
			return err
		}
	}
	return nil
}

func NewCustomerOrderUseCase(billRepo repository.BillRepository, tableRepo repository.TableRepository) CustomerOrderUseCase {
	usecase := new(customerOrderUseCase)
	usecase.billRepo = billRepo
	usecase.tableRepo = tableRepo
	return usecase
}
