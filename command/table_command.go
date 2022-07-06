package command

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	usecase "livecode-wmb-rest-api/usecase/table"
	"livecode-wmb-rest-api/utils"
	"log"
)

func TableCreate(tableRepo repository.TableRepository) {
	// Create (C)
	CreateTableUseCase := usecase.NewCreateTableUseCase(tableRepo)
	table01 := model.Table{
		TableDescription: "T001",
	}
	table02 := model.Table{
		TableDescription: "T002",
	}
	table03 := model.Table{
		TableDescription: "T003",
	}
	err := CreateTableUseCase.CreateTable([]*model.Table{&table01, &table02, &table03})
	utils.RaiseError(err)
}

func TableRead(tableRepo repository.TableRepository) {
	// Read (R)
	ReadTableUseCase := usecase.NewReadTableUseCase(tableRepo)
	tableAll, err := ReadTableUseCase.ReadAllTable()
	utils.RaiseError(err)
	log.Println(tableAll)

	tablewithID, err := ReadTableUseCase.ReadTableById(3)
	utils.RaiseError(err)
	log.Println(tablewithID)

	tablewithName, err := ReadTableUseCase.ReadTableBy(map[string]interface{}{"table_name": "Nasi Uduk"})
	utils.RaiseError(err)
	log.Println(tablewithName)
}

func TableUpdate(tableRepo repository.TableRepository) {
	// Cari table yang ingin di update
	ReadTableUseCase := usecase.NewReadTableUseCase(tableRepo)
	tablewithID, err := ReadTableUseCase.ReadTableById(3)
	utils.RaiseError(err)
	log.Println(tablewithID)

	// Misal ganti nama
	UpdateTableUseCase := usecase.NewUpdateTableUseCase(tableRepo)
	UpdateTableUseCase.UpdateTable(&tablewithID, map[string]interface{}{"table_name": "Nasi Liwet"})

	// Dicetak kembali untuk dicek apakah sudah berubah
	tablewithID, err = ReadTableUseCase.ReadTableById(3)
	utils.RaiseError(err)
	log.Println(tablewithID)
}

func TableDelete(tableRepo repository.TableRepository) {
	// Cari table yang ingin di update
	ReadTableUseCase := usecase.NewReadTableUseCase(tableRepo)
	tablewithID, err := ReadTableUseCase.ReadTableById(3)
	utils.RaiseError(err)
	log.Println(tablewithID)

	// Proses delete dilakukan
	DeleteTableUseCase := usecase.NewDeleteTableUseCase(tableRepo)
	DeleteTableUseCase.DeleteTable(&tablewithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	tablewithID, err = ReadTableUseCase.ReadTableById(3)
	utils.RaiseError(err)
	log.Println(tablewithID)
}
