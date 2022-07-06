package manager

import (
	"livecode-wmb-rest-api/config"
	"livecode-wmb-rest-api/migration"
	"livecode-wmb-rest-api/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Infra disini bertugas sebagai database penyimpanan pengganti slice
type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (dbReturn *gorm.DB, err error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	utils.PanicError(err)

	env := os.Getenv("DB_VENV")
	if env == "migration" {
		log.Println("Running in Migration Environment")
		dbReturn = db.Debug()
		migration.MigrateDb(dbReturn)
	} else if env == "dev" {
		log.Println("Running in Development Environment")
		dbReturn = db.Debug()
	}
	return
}
