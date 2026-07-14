package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT  string
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
		PORT:  getEnv("PORT", "8080"),
		DBURL: getEnv("DATABASE_URL", "postgresql://postgres:password@localhost:5432"),
	}
}
