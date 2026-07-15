package main

import (
	"fmt"
	"grpc/internal/config"
	"grpc/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := config.LoadConfig()

	db := database.Connect(cfg.DBURL)

	fmt.Print(db)

	r.Run(cfg.PORT)
}
