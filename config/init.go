package config

import (
	"lopmartyn-gin-swagger/config/database"
	"os"
)

type GoDotENV struct {
	Port           string                  `json:"port"`
	GinMode        string                  `json:"ginMode"`
	PostgresConfig database.PostgresConfig `json:"postgres"`
}

// GetGoDotENV Get .env configuration
func GetGoDotENV() GoDotENV {
	return GoDotENV{
		Port:           ":" + os.Getenv("PORT"),
		GinMode:        os.Getenv("GIN_MODE"),
		PostgresConfig: setPostgres(),
	}
}

// setPostgres Get database configuration for bunDB
func setPostgres() database.PostgresConfig {
	return database.PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DBNAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}
}
