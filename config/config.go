package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ApiConfig struct {
	DatabaseDSN  string
	JwtSecretKey string
}

var cfg *ApiConfig

func LoadApiConfig() (*ApiConfig, error) {
	if os.Getenv("ENVIRONMENT") == "" {
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}
	cfg = &ApiConfig{
		DatabaseDSN:  os.Getenv("DATABASE_DSN"),
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),
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
