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