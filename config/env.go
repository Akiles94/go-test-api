package config

import (
	"log"
	"os"
	"strconv"

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
	rateLimitCount, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_COUNT"))

	Env = &EnvConfig{
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		ApiPort:        os.Getenv("API_PORT"),
		Mode:           os.Getenv("MODE"),
		RateLimitCount: rateLimitCount,
		JWTSecret:      os.Getenv("JWT_SECRET"),
		JWTRefreshExpirationDays: func() int {
			val, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION_DAYS"))
			return val
		}(),
	}
}
