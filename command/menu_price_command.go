package command

import (
	"database/sql"
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	usecase "livecode-wmb-rest-api/usecase/menu_price"
	"livecode-wmb-rest-api/utils"
	"log"
)

func MenuPriceCreate(menuPriceRepo repository.MenuPriceRepository) {
	// Create (C)
	CreateMenuPriceUseCase := usecase.NewCreateMenuPriceUseCase(menuPriceRepo)
	menuPrice01 := model.MenuPrice{
		Price: 2000,
	}
	menuPrice02 := model.MenuPrice{
		Price:  5000,
		MenuID: sql.NullInt64{Int64: 2, Valid: true},
	}
	menuPrice03 := model.MenuPrice{
		Price:  10000,
		MenuID: sql.NullInt64{Int64: 3, Valid: true},
	}
	err := CreateMenuPriceUseCase.CreateMenuPrice([]*model.MenuPrice{&menuPrice01, &menuPrice02, &menuPrice03})
	utils.RaiseError(err)
}

func MenuPriceRead(menuPriceRepo repository.MenuPriceRepository) {
	// Read (R)
	ReadMenuPriceUseCase := usecase.NewReadMenuPriceUseCase(menuPriceRepo)
	menuPriceAll, err := ReadMenuPriceUseCase.ReadAllMenuPrice()
	utils.RaiseError(err)
	log.Println(menuPriceAll)

	menuPricewithID, err := ReadMenuPriceUseCase.ReadMenuPriceById(3)
	utils.RaiseError(err)
	log.Println(menuPricewithID)

	menuPricewithName, err := ReadMenuPriceUseCase.ReadMenuPriceBy(map[string]interface{}{"menuPrice_name": "Nasi Uduk"})
	utils.RaiseError(err)
	log.Println(menuPricewithName)
}

func MenuPriceUpdate(menuPriceRepo repository.MenuPriceRepository) {
	// Cari menuPrice yang ingin di update
	ReadMenuPriceUseCase := usecase.NewReadMenuPriceUseCase(menuPriceRepo)
	menuPricewithID, err := ReadMenuPriceUseCase.ReadMenuPriceById(3)
	utils.RaiseError(err)
	log.Println(menuPricewithID)

	// Misal ganti nama
	UpdateMenuPriceUseCase := usecase.NewUpdateMenuPriceUseCase(menuPriceRepo)
	UpdateMenuPriceUseCase.UpdateMenuPrice(&menuPricewithID, map[string]interface{}{"menuPrice_name": "Nasi Liwet"})

	// Dicetak kembali untuk dicek apakah sudah berubah
	menuPricewithID, err = ReadMenuPriceUseCase.ReadMenuPriceById(3)
	utils.RaiseError(err)
	log.Println(menuPricewithID)
}

func MenuPriceDelete(menuPriceRepo repository.MenuPriceRepository) {
	// Cari menuPrice yang ingin di update
	ReadMenuPriceUseCase := usecase.NewReadMenuPriceUseCase(menuPriceRepo)
	menuPricewithID, err := ReadMenuPriceUseCase.ReadMenuPriceById(1)
	utils.RaiseError(err)
	log.Println(menuPricewithID)

	// Proses delete dilakukan
	DeleteMenuPriceUseCase := usecase.NewDeleteMenuPriceUseCase(menuPriceRepo)
	DeleteMenuPriceUseCase.DeleteMenuPrice(&menuPricewithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	menuPricewithID, err = ReadMenuPriceUseCase.ReadMenuPriceById(1)
	utils.RaiseError(err)
	log.Println(menuPricewithID)
}
