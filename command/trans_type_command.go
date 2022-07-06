package command

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/repository"
	usecase "livecode-wmb-rest-api/usecase/trans_type"
	"livecode-wmb-rest-api/utils"
	"log"
)

func TransTypeCreate(transTypeRepo repository.TransTypeRepository) {
	// Create (C)
	CreateTransTypeUseCase := usecase.NewCreateTransTypeUseCase(transTypeRepo)
	transType01 := model.TransType{
		ID:          "EI",
		Description: "Eat In",
	}
	transType02 := model.TransType{
		ID:          "TA",
		Description: "Take Away",
	}
	err := CreateTransTypeUseCase.CreateTransType([]*model.TransType{&transType01, &transType02})
	utils.RaiseError(err)
}

func TransTypeRead(transTypeRepo repository.TransTypeRepository) {
	// Read (R)
	ReadTransTypeUseCase := usecase.NewReadTransTypeUseCase(transTypeRepo)
	transTypeAll, err := ReadTransTypeUseCase.ReadAllTransType()
	utils.RaiseError(err)
	log.Println(transTypeAll)

	transTypewithID, err := ReadTransTypeUseCase.ReadTransTypeById(3)
	utils.RaiseError(err)
	log.Println(transTypewithID)

	transTypewithName, err := ReadTransTypeUseCase.ReadTransTypeBy(map[string]interface{}{"transType_name": "Nasi Uduk"})
	utils.RaiseError(err)
	log.Println(transTypewithName)
}

func TransTypeUpdate(transTypeRepo repository.TransTypeRepository) {
	// Cari transType yang ingin di update
	ReadTransTypeUseCase := usecase.NewReadTransTypeUseCase(transTypeRepo)
	transTypewithID, err := ReadTransTypeUseCase.ReadTransTypeById(3)
	utils.RaiseError(err)
	log.Println(transTypewithID)

	// Misal ganti nama
	UpdateTransTypeUseCase := usecase.NewUpdateTransTypeUseCase(transTypeRepo)
	UpdateTransTypeUseCase.UpdateTransType(&transTypewithID, map[string]interface{}{"transType_name": "Nasi Liwet"})

	// Dicetak kembali untuk dicek apakah sudah berubah
	transTypewithID, err = ReadTransTypeUseCase.ReadTransTypeById(3)
	utils.RaiseError(err)
	log.Println(transTypewithID)
}

func TransTypeDelete(transTypeRepo repository.TransTypeRepository) {
	// Cari transType yang ingin di update
	ReadTransTypeUseCase := usecase.NewReadTransTypeUseCase(transTypeRepo)
	transTypewithID, err := ReadTransTypeUseCase.ReadTransTypeById(3)
	utils.RaiseError(err)
	log.Println(transTypewithID)

	// Proses delete dilakukan
	DeleteTransTypeUseCase := usecase.NewDeleteTransTypeUseCase(transTypeRepo)
	DeleteTransTypeUseCase.DeleteTransType(&transTypewithID)

	// Dicetak kembali untuk dicek apakah sudah berubah
	transTypewithID, err = ReadTransTypeUseCase.ReadTransTypeById(3)
	utils.RaiseError(err)
	log.Println(transTypewithID)
}
