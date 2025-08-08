package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBHost                   string
	DBPort                   string
	DBUser                   string
	DBPassword               string
	DBName                   string
	ApiPort                  string
	Mode                     string
	RateLimitCount           int
	JWTSecret                string
	JWTRefreshExpirationDays int
}

var Env *EnvConfig

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}

	Env = &EnvConfig{
		DBHost:     os.Getenv("PRODUCT_DB_HOST"),
		DBPort:     os.Getenv("PRODUCT_DB_PORT"),
		DBUser:     os.Getenv("PRODUCT_DB_USER"),
		DBPassword: os.Getenv("PRODUCT_DB_PASSWORD"),
		DBName:     os.Getenv("PRODUCT_DB_NAME"),
		ApiPort:    os.Getenv("PRODUCT_API_PORT"),
		Mode:       os.Getenv("MODE"),
	}
}
