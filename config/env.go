package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var Env *EnvConfig

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}

	Env = &EnvConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
