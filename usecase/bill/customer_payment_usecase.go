package bill_usecase

import (
	"errors"
	"fmt"
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
	"livecode-wmb-2/utils"
)

type CustomerPaymentUseCase interface {
	OrderPayment(billId uint, paymentMethod string) (model.Bill, error)
	PrintBill(billId uint) (model.Bill, error)
}

type customerPaymentUseCase struct {
	billRepo  repository.BillRepository
	tableRepo repository.TableRepository
}

func (c *customerPaymentUseCase) PrintBill(billId uint) (model.Bill, error) {
	billSlice, err := c.billRepo.FindBy(map[string]interface{}{"id": billId}, []string{"BillDetails", "BillDetails.MenuPrice", "Customer", "Customer.Discounts"})
	if err != nil {
		return model.Bill{}, utils.DataNotFoundError()
	}
	return billSlice[0], nil
}

func (c *customerPaymentUseCase) OrderPayment(billId uint, paymentMethod string) (model.Bill, error) {
	if paymentMethod != "tunai" && paymentMethod != "lopei" {
		return model.Bill{}, errors.New("metode pembayaran yang dipilih tidak sesuai")
	}
	billSlice, err := c.billRepo.FindBy(map[string]interface{}{"id": billId}, []string{"BillDetails", "BillDetails.MenuPrice", "Customer", "Customer.Discounts"})
	if err != nil || len(billSlice) == 0 {
		return model.Bill{}, utils.DataNotFoundError()
	}
	if paymentMethod == "lopei" {
		var grandTotal float32
		fmt.Println(billSlice[0].BillDetails)
		for _, billDetail := range billSlice[0].BillDetails {
			grandTotal += billDetail.MenuPrice.Price*billDetail.Qty
		}
		if billSlice[0].Customer.IsMember {
			grandTotal = grandTotal*float32(100-billSlice[0].Customer.Discounts[0].Pct)
			fmt.Println(grandTotal)
		}
		if err := c.billRepo.LopeiPayment(int32(billSlice[0].Customer.ID), grandTotal); err != nil {
			return model.Bill{}, err
		}
	}
	if billSlice[0].TransTypeID == "EI" {
		tableSlice, err := c.tableRepo.FindBy(map[string]interface{}{"id": billSlice[0].TableID})
		if err != nil {
			fmt.Println("Informasi meja dalam bill tidak ditemukan.")
			return model.Bill{}, utils.DataNotFoundError()
		}
		tableSelected := tableSlice[0]
		// Tidak perlu transactional karena yang diupdate hanya table
		err = c.tableRepo.UpdateBy(&tableSelected, map[string]interface{}{"is_available": true})
		if err != nil {
			fmt.Println("Update informasi meja gagal dilakukan.")
			return model.Bill{}, err
		}
	}
	return billSlice[0], nil
}

func NewCustomerPaymentUseCase(billRepo repository.BillRepository, tableRepo repository.TableRepository) CustomerPaymentUseCase {
	usecase := new(customerPaymentUseCase)
	usecase.billRepo = billRepo
	usecase.tableRepo = tableRepo
	return usecase
}
