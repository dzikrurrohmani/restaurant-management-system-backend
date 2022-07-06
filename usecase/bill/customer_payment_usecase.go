package bill_usecase

import (
	"fmt"
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
)

type CustomerPaymentUseCase interface {
	OrderPayment(billId uint) (model.Bill, error)
	PrintBill(billId uint) (model.Bill, error)
}

type customerPaymentUseCase struct {
	billRepo  repository.BillRepository
	tableRepo repository.TableRepository
}

func (c *customerPaymentUseCase) PrintBill(billId uint) (model.Bill, error) {
	billSlice, err := c.billRepo.FindBy(map[string]interface{}{"id": billId}, []string{"BillDetails", "BillDetails.MenuPrice"})
	if err != nil {
		fmt.Println("Informasi bill tidak ditemukan.")
		return model.Bill{}, err
	}
	return billSlice[0], nil
}

func (c *customerPaymentUseCase) OrderPayment(billId uint) (model.Bill, error) {
	billSlice, err := c.billRepo.FindBy(map[string]interface{}{"id": billId}, nil)
	if err != nil && len(billSlice)==0 {
		fmt.Println("Informasi bill tidak ditemukan.")
		return model.Bill{}, err
	}
	if billSlice[0].TransTypeID == "EI" {
		tableSlice, err := c.tableRepo.FindBy(map[string]interface{}{"id": billSlice[0].TableID})
		if err != nil {
			fmt.Println("Informasi meja dalam bill tidak ditemukan.")
			return model.Bill{}, err
		}
		tableSelected := tableSlice[0]
		// Tidak perlu transactional karena yang diupdate hanya table
		err = c.tableRepo.UpdateBy(&tableSelected, map[string]interface{}{"is_available": true})
		if err != nil {
			fmt.Println("Update informasi meja gagal dilakukan.")
			return model.Bill{}, err
		}
	}
	return c.PrintBill(billId)
}

func NewCustomerPaymentUseCase(billRepo repository.BillRepository, tableRepo repository.TableRepository) CustomerPaymentUseCase {
	usecase := new(customerPaymentUseCase)
	usecase.billRepo = billRepo
	usecase.tableRepo = tableRepo
	return usecase
}
