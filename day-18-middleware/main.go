package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// logging middleware

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		fmt.Println("Completed")
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := r.Header.Get("Authorization")

		if secret != "secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			ID:   1,
			Name: "Tony",
			Age:  12,
		},
		{
			ID:   2,
			Name: "Stark",
			Age:  13,
		},
		{
			ID:   3,
			Name: "Banner",
			Age:  14,
		},
	}

	w.Header().Set("Content-Type", "applicaton/json")

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func main() {

	r := chi.NewRouter()

	r.Use(loggingMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Hello World</h1>")
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/users", getUsers)
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
