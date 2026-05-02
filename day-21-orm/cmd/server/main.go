package main

import (
	"net/http"
	"orm/db"
	"orm/handlers"
	"orm/repository"
	"orm/routes"

	"github.com/go-chi/chi/v5"
)

func main() {
	database := db.ConnectDB()

	repo := &repository.UserRepository{DB: database}
	h := &handlers.UserHandler{Repo: repo}

	r := chi.NewRouter()
	routes.RegsiterRouter(h, r)

	http.ListenAndServe(":8080", r)
}
