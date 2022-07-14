package config

import (
	"fmt"
	"os"
)

type GrpcConfig struct {
	GrpcUrl string
}

type ApiConfig struct {
	ApiUrl string
}

type DbConfig struct {
	DataSourceName string
}

type Config struct {
	ApiConfig
	DbConfig
	GrpcConfig
}

func (c *Config) readConfig() {
	apiUrl := os.Getenv("API_URL")
	grpcUrl := os.Getenv("GRPC_URL")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	c.DbConfig = DbConfig{DataSourceName: dsn}
	c.ApiConfig = ApiConfig{ApiUrl: apiUrl}
	c.GrpcConfig = GrpcConfig{GrpcUrl: grpcUrl}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
