package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	connStr := os.Getenv("DATABASE_URL")

	fmt.Println("Connection String: ", connStr)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database", err.Error())
	}

	fmt.Println("Database connected successfully")

	return db
}
