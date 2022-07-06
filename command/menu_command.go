package command

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	usecase "livecode-wmb-rest-api/usecase/menu"
	"livecode-wmb-rest-api/utils"
	"log"
)

func MenuCreate(menuRepo repository.MenuRepository) {
	// Create (C)
	CreateMenuUseCase := usecase.NewCreateMenuUseCase(menuRepo)
	menu01 := model.Menu{
		MenuName: "Nasi Goreng",
	}
	menu02 := model.Menu{
		MenuName: "Nasi Padang",
	}
	menu03 := model.Menu{
		MenuName: "Nasi Uduk",
	}
	err := CreateMenuUseCase.CreateMenu([]*model.Menu{&menu01, &menu02, &menu03})
	utils.RaiseError(err)
}

func MenuRead(menuRepo repository.MenuRepository) {
	// Read (R)
	ReadMenuUseCase := usecase.NewReadMenuUseCase(menuRepo)
	menuAll, err := ReadMenuUseCase.ReadAllMenu()
	utils.RaiseError(err)
	log.Println(menuAll)

	menuwithID, err := ReadMenuUseCase.ReadMenuById(3)
	utils.RaiseError(err)
	log.Println(menuwithID)

	menuwithName, err := ReadMenuUseCase.ReadMenuBy(map[string]interface{}{"menu_name": "Nasi Uduk"})
	utils.RaiseError(err)
	log.Println(menuwithName)
}

func MenuUpdate(menuRepo repository.MenuRepository) {
	// Cari menu yang ingin di update
	ReadMenuUseCase := usecase.NewReadMenuUseCase(menuRepo)
	menuwithID, err := ReadMenuUseCase.ReadMenuById(3)
	utils.RaiseError(err)
	log.Println(menuwithID)

	// Misal ganti nama
	UpdateMenuUseCase := usecase.NewUpdateMenuUseCase(menuRepo)
	UpdateMenuUseCase.UpdateMenu(&menuwithID, map[string]interface{}{"menu_name": "Nasi Liwet"})

	// Dicetak kembali untuk dicek apakah sudah berubah
	menuwithID, err = ReadMenuUseCase.ReadMenuById(3)
	utils.RaiseError(err)
	log.Println(menuwithID)
}

func MenuDelete(menuRepo repository.MenuRepository) {
	// Cari menu yang ingin di update
	ReadMenuUseCase := usecase.NewReadMenuUseCase(menuRepo)
	menuwithID, err := ReadMenuUseCase.ReadMenuById(3)
	utils.RaiseError(err)
	log.Println(menuwithID)

	// Proses delete dilakukan
	DeleteMenuUseCase := usecase.NewDeleteMenuUseCase(menuRepo)
	DeleteMenuUseCase.DeleteMenu(&menuwithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	menuwithID, err = ReadMenuUseCase.ReadMenuById(3)
	utils.RaiseError(err)
	log.Println(menuwithID)
}
