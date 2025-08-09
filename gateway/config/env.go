package config

import (
	"log"
	"os"
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
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}
	rateLimitCount, _ := strconv.Atoi(os.Getenv("RATE_LIMIT_COUNT"))
	jwtRefreshExpirationDays, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION_DAYS"))

	Env = &EnvConfig{
		Mode:                     os.Getenv("MODE"),
		RateLimitCount:           rateLimitCount,
		JWTSecret:                os.Getenv("JWT_SECRET"),
		JWTRefreshExpirationDays: jwtRefreshExpirationDays,
		GRPCPort:                 os.Getenv("GRPC_PORT"),
		GRPCHost:                 os.Getenv("GRPC_HOST"),
	}
}
