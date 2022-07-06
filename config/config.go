package config

import (
	"fmt"
	"os"
)

type ApiConfig struct {
	Url string
}

type DbConfig struct {
	DataSourceName string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.ApiConfig = ApiConfig{Url: api}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}


// type dBConfig struct {
// 	dbVenv string
// 	dbHost string
// 	dbPort string
// 	dbUser string
// 	dbPass string
// 	dbName string
// }


// func (c *Config) initDb() {
// 	dbVenv := os.Getenv("DB_ENV")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")

// 	dBConfig := dBConfig{
// 		dbVenv, dbHost, dbPort, dbUser, dbPass, dbName,
// 	}

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dBConfig.dbHost, dBConfig.dbUser, dBConfig.dbPass, dBConfig.dbName, dBConfig.dbPort)

// 	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	util.PanicError(err)

// 	c.checkConnection(gormDB)

// 	if dBConfig.dbVenv == "development" {
// 		log.Println("Running in Development Environment")
// 		c.db = gormDB.Debug()
// 	} else if dBConfig.dbVenv == "production" {
// 		log.Println("Running in Production Environment")
// 		c.db = gormDB
// 	} else {
// 		util.PanicError(errors.New(fmt.Sprintf("%s in wrong environment",dBConfig.dbVenv)))
// 	}
// }

// func (c *Config) checkConnection (gormDB *gorm.DB) {
// 	runningDb, err := gormDB.DB()
// 	util.PanicError(err)
// 	err = runningDb.Ping()
// 	util.PanicError(err)
// }

// func (c *Config) DbConn() *gorm.DB {
// 	return c.db
// }

// func (c *Config) DbClose() {
// 	runningDb, err := c.db.DB()
// 	util.PanicError(err)
// 	defer func(runningDb *sql.DB) {
// 		err := runningDb.Close()
// 		util.PanicError(err)
// 	}(runningDb)
// }
// func NewConfigDB() *Config {
// 	cfg := Config{}
// 	cfg.initDb()
// 	return &cfg
// }
