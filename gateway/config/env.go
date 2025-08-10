package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ApiPort                  string
	Mode                     string
	RateLimitCount           int
	JWTSecret                string
	JWTRefreshExpirationDays int
	GRPCHost                 string
	GRPCPort                 string
}

var Env *EnvConfig

func LoadEnv() {
	envPath := filepath.Join("gateway", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}
	rateLimitCount, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_COUNT"))
	jwtRefreshExpirationDays, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION_DAYS"))

	Env = &EnvConfig{
		ApiPort:                  os.Getenv("API_PORT"),
		Mode:                     os.Getenv("MODE"),
		RateLimitCount:           rateLimitCount,
		JWTSecret:                os.Getenv("JWT_SECRET"),
		JWTRefreshExpirationDays: jwtRefreshExpirationDays,
		GRPCPort:                 os.Getenv("GRPC_PORT"),
		GRPCHost:                 os.Getenv("GRPC_HOST"),
	}
}
