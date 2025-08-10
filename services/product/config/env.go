package config

import (
	"log"
	"os"
	"path/filepath"

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
	GatewayGRPCAddress       string
	ServiceHost              string
}

var Env *EnvConfig

func LoadEnv() {
	envPath := filepath.Join("services/product", ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Println("⚠️  No .env file found, using system env variables")
	}

	Env = &EnvConfig{
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		ApiPort:            os.Getenv("API_PORT"),
		Mode:               os.Getenv("MODE"),
		GatewayGRPCAddress: os.Getenv("GATEWAY_GRPC_ADDRESS"),
		ServiceHost:        os.Getenv("SERVICE_HOST"),
	}
}
