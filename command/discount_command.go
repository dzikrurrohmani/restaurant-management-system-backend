package command

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	usecase "livecode-wmb-rest-api/usecase/discount"
	"livecode-wmb-rest-api/utils"
	"log"
)

func DiscountCreate(discountRepo repository.DiscountRepository) {
	// Create (C)
	CreateDiscountUseCase := usecase.NewCreateDiscountUseCase(discountRepo)
	discount01 := model.Discount{
		Pct:         5,
		Description: "Mendapat potongan harga sejumlah 5 persen",
	}
	discount02 := model.Discount{
		Pct:         7,
		Description: "Mendapat potongan harga sejumlah 7 persen",
	}
	discount03 := model.Discount{
		Pct:         10,
		Description: "Mendapat potongan harga sejumlah 10 persen",
	}
	err := CreateDiscountUseCase.CreateDiscount([]*model.Discount{&discount01, &discount02, &discount03})
	utils.RaiseError(err)
}

func DiscountRead(discountRepo repository.DiscountRepository) {
	// Read (R)
	ReadDiscountUseCase := usecase.NewReadDiscountUseCase(discountRepo)
	discountAll, err := ReadDiscountUseCase.ReadAllDiscount()
	utils.RaiseError(err)
	log.Println(discountAll)

	discountwithID, err := ReadDiscountUseCase.ReadDiscountById(3)
	utils.RaiseError(err)
	log.Println(discountwithID)

	discountwithName, err := ReadDiscountUseCase.ReadDiscountBy(map[string]interface{}{"discount_name": "Nasi Uduk"})
	utils.RaiseError(err)
	log.Println(discountwithName)
}

func DiscountUpdate(discountRepo repository.DiscountRepository) {
	// Cari discount yang ingin di update
	ReadDiscountUseCase := usecase.NewReadDiscountUseCase(discountRepo)
	discountwithID, err := ReadDiscountUseCase.ReadDiscountById(3)
	utils.RaiseError(err)
	log.Println(discountwithID)

	// Misal ganti nama
	UpdateDiscountUseCase := usecase.NewUpdateDiscountUseCase(discountRepo)
	UpdateDiscountUseCase.UpdateDiscount(&discountwithID, map[string]interface{}{"discount_name": "Nasi Liwet"})

	// Dicetak kembali untuk dicek apakah sudah berubah
	discountwithID, err = ReadDiscountUseCase.ReadDiscountById(3)
	utils.RaiseError(err)
	log.Println(discountwithID)
}

func DiscountDelete(discountRepo repository.DiscountRepository) {
	// Cari discount yang ingin di update
	ReadDiscountUseCase := usecase.NewReadDiscountUseCase(discountRepo)
	discountwithID, err := ReadDiscountUseCase.ReadDiscountById(3)
	utils.RaiseError(err)
	log.Println(discountwithID)

	// Proses delete dilakukan
	DeleteDiscountUseCase := usecase.NewDeleteDiscountUseCase(discountRepo)
	DeleteDiscountUseCase.DeleteDiscount(&discountwithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	discountwithID, err = ReadDiscountUseCase.ReadDiscountById(3)
	utils.RaiseError(err)
	log.Println(discountwithID)
}
