package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBURL string
	ENV   string
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}

func LoadConfig() Config {
	// load .env file for local development
	_ = godotenv.Load()

	return Config{
		Port:  getEnv("PORT", "8080"),
		DBURL: getEnv("DB_URL", ""),
		ENV:   getEnv("ENV", "development"),
	}
}
