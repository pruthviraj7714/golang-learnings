package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURL string
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)

	if val == "" {
		return fallback
	}

	return val
}

func LoadConfig() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error while loading env")
	}

	return &Config{
		DBURL: getEnv("DATABASE_URL", "postgresql://postgres:password@localhost:5432"),
	}
}
