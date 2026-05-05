package main

import (
	"config/config"
	"config/db"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	db.Connect(cfg.DBURL)

	fmt.Println("Server running on port", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
