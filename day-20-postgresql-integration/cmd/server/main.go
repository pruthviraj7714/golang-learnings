package main

import (
	"log"
	"net/http"
	"postgresql-integration/db"
	"postgresql-integration/handlers"
	"postgresql-integration/repository"
	"postgresql-integration/routes"

	_ "github.com/lib/pq"
)

func main() {
	database := db.Connect()
	defer database.Close()

	repo := &repository.UserRepository{DB: database}
	handler := &handlers.UserHandler{Repo: repo}

	routes.RegisterRouter(handler)

	log.Println("Server is running on http://localhost:8080 🚀")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
