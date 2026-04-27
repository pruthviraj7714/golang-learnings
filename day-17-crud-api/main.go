package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = make(map[string]User)

func getAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	var result []User

	for _, user := range users {
		result = append(result, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	users[user.ID] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, ok := users[id]

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updatedUser User

	err := json.NewDecoder(r.Body).Decode(&updatedUser)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	updatedUser.ID = id
	users[id] = updatedUser

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_, ok := users[id]

	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, id)

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", getAllUsersHandler)
		r.Post("/users", createUserHandler)
		r.Get("/users/{id}", getUserByIDHandler)
		r.Put("/users/{id}", updateUserHandler)
		r.Delete("/users/{id}", deleteUserByIDHandler)
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	fmt.Println("Server is running on port 8080")
}
