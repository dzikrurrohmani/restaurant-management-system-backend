package manager

import (
	"livecode-wmb-2/config"
	"livecode-wmb-2/migration"
	"livecode-wmb-2/service"
	"livecode-wmb-2/utils"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Infra disini bertugas sebagai database penyimpanan pengganti slice
type Infra interface {
	LopeiClientConn() service.LopeiPaymentClient
	SqlDb() *gorm.DB
}

type infra struct {
	lopeiClient service.LopeiPaymentClient
	db          *gorm.DB
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func initDbResource(dataSourceName string) (dbReturn *gorm.DB, err error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	utils.PanicError(err)

	env := os.Getenv("DB_VENV")
	if env == "migration" {
		log.Println("Running in Migration Environment")
		dbReturn = db.Debug()
		migration.MigrateDb(dbReturn)
	} else if env == "development" {
		log.Println("Running in Development Environment")
		dbReturn = db.Debug()
	}
	return
}

func (i *infra) LopeiClientConn() service.LopeiPaymentClient {
	return i.lopeiClient
}

func initGrpcClient(grpcUrl string) service.LopeiPaymentClient {
	dial, err := grpc.Dial(grpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect..", err)
	}
	client := service.NewLopeiPaymentClient(dial)
	log.Println("GRPC client connected..")
	return client
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := initGrpcClient(config.GrpcUrl)
	return &infra{db: resource, lopeiClient: client}
}
