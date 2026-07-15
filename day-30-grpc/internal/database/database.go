package database

import (
	"grpc/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(connString string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
