package bill_usecase

import (
	"fmt"
	"livecode-wmb-rest-api/repository"
	"log"
)

type CustomerPaymentUseCase interface {
	OrderPayment(billId uint) error
	PrintBill(billId uint) error
}

type customerPaymentUseCase struct {
	billRepo  repository.BillRepository
	tableRepo repository.TableRepository
}

func (c *customerPaymentUseCase) PrintBill(billId uint) error {
	billSlice, err := c.billRepo.FindBy(map[string]interface{}{"id": billId}, []string{"BillDetails", "BillDetails.MenuPrice"})
	if err != nil {
		fmt.Println("Informasi bill tidak ditemukan.")
		return err
	}
	log.Println(billSlice[0].ToString())
	return nil
}

func (c *customerPaymentUseCase) OrderPayment(billId uint) error {
	billSlice, err := c.billRepo.FindBy(map[string]interface{}{"id": billId}, nil)
	if err != nil {
		fmt.Println("Informasi bill tidak ditemukan.")
		return err
	}
	tableSlice, err := c.tableRepo.FindBy(map[string]interface{}{"id": billSlice[0].TableID})
	if err != nil {
		fmt.Println("Informasi meja dalam bill tidak ditemukan.")
		return err
	}
	tableSelected := tableSlice[0]
	// Tidak perlu transactional karena yang diupdate hanya table
	err = c.tableRepo.UpdateBy(&tableSelected, map[string]interface{}{"is_available": true})
	if err != nil {
		fmt.Println("Update informasi meja gagal dilakukan.")
		return err
	}
	c.PrintBill(billId)
	return nil
}

func NewCustomerPaymentUseCase(billRepo repository.BillRepository, tableRepo repository.TableRepository) CustomerPaymentUseCase {
	usecase := new(customerPaymentUseCase)
	usecase.billRepo = billRepo
	usecase.tableRepo = tableRepo
	return usecase
}
