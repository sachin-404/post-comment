package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ApiConfig struct {
	DatabaseDSN string
}

var cfg *ApiConfig

func LoadApiConfig() (*ApiConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	cfg := &ApiConfig{
		DatabaseDSN: os.Getenv("DATABASE_DSN"),
	}
	return cfg, nil
}

func GetApiConfig() *ApiConfig {
	if cfg == nil {
		if _, err := LoadApiConfig(); err != nil {
			log.Fatalf("failed to load config: %v", err.Error())
		}
	}
	return cfg
}
