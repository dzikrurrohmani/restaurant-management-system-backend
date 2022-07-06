package migration

import (
	"livecode-wmb-rest-api/model"
	"livecode-wmb-rest-api/utils"

	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Customer{},
		&model.Menu{},
		&model.Table{},
		&model.Discount{},
		&model.MenuPrice{},
		&model.TransType{},
		&model.BillDetail{},
		&model.Bill{},
	)
	utils.PanicError(err)
}
