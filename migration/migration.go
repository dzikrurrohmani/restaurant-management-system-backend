package migration

import (
	"livecode-wmb-2/model"
	"livecode-wmb-2/utils"

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
