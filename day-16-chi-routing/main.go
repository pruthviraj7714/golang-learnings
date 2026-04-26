package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username != "" {
		fmt.Fprintf(w, "Hello, %s", username)
	} else {
		fmt.Fprint(w, "Hello, Guest")
	}
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	u := User{
		Name: "Tony",
		Age:  26,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(u)
}

func getUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	fmt.Fprintf(w, "User Id: %s", userId)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from GO API 🚀")
	})

	r.Get("/hello", helloHandler)
	r.Get("/user", userHandler)

	//route grouping
	r.Route("/api/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Users list")
		})
		r.Get("/{id}", getUserByIdHandler)
	})

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
