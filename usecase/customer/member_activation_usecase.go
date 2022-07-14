package customer_usecase

import (
	"errors"
	"fmt"
	"livecode-wmb-2/model"
	"livecode-wmb-2/repository"
)

type MemberActivationUseCase interface {
	ActivateMember(phoneNumber string, discountId uint) (model.Customer, error)
}

type memberActivationUseCase struct {
	custRepo     repository.CustomerRepository
	discountRepo repository.DiscountRepository
}

func (m *memberActivationUseCase) ActivateMember(phoneNumber string, discountId uint) (model.Customer, error) {
	by := map[string]interface{}{"mobile_phone_no": phoneNumber}
	customerSlice, err := m.custRepo.FindBy(by)
	if err != nil {
		fmt.Println("Data tidak ditemukan.")
		return model.Customer{}, err
	}
	customerSelected := customerSlice[0]
	if customerSelected.IsMember {
		fmt.Println("Data tidak ditemukan.")
		return customerSelected, errors.New("nomor telepon tersebut sudah terdaftar sebagai member")
	}
	err = m.custRepo.UpdateBy(&customerSelected, map[string]interface{}{"is_member": true})
	if err != nil {
		fmt.Println("Aktivasi member gagal dilakukan.")
		return model.Customer{}, err
	}
	fmt.Println("Aktivasi member berhasil.")
	// Belum selesai tambah discount
	discountSlice, _ := m.discountRepo.FindBy(map[string]interface{}{"id": discountId})
	err = m.custRepo.UpdateAssociation(&customerSelected, "Discounts", &discountSlice[0])
	if err != nil {
		fmt.Println("gagal menambahkan previlledge discount")
	}
	customerSlice, err = m.custRepo.FindBy(by)
	if err != nil {
		fmt.Println("Data tidak ditemukan.")
		return model.Customer{}, err
	}
	customerReturn := customerSlice[0]
	return customerReturn, nil
}

func NewMemberActivationUseCase(custRepo repository.CustomerRepository, discountRepo repository.DiscountRepository) MemberActivationUseCase {
	usecase := new(memberActivationUseCase)
	usecase.custRepo = custRepo
	usecase.discountRepo = discountRepo
	return usecase
}
