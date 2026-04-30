package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var users = make(map[string]User)

func writeJSONError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	decoder := json.NewDecoder(r.Body)
	// This tells the decoder to reject JSON objects with extra fields
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&user); err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if user.ID == "" || user.Name == "" {
		writeJSONError(w, "ID and Name are required", http.StatusBadRequest)
		return
	}

	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var result []User

	for _, user := range users {
		result = append(result, user)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func deleteUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if _, ok := users[id]; !ok {
		writeJSONError(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, id)

	w.WriteHeader(http.StatusNoContent)
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, ok := users[id]

	if !ok {
		writeJSONError(w, "User not fouhnd", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

type UpdateUser struct {
	//Pointer allows:
	// 1. detect missing vs empty
	// 2. ignore missing fields
	Name *string `json:"name"`
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, ok := users[id]

	if !ok {
		writeJSONError(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser UpdateUser

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&updatedUser); err != nil {
		writeJSONError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if decoder.More() {
		writeJSONError(w, "Only one JSON object is allowed", http.StatusBadRequest)
		return
	}

	if updatedUser.Name == nil {
		writeJSONError(w, "At least one field is required: 'name'", http.StatusBadRequest)
		return
	}

	if updatedUser.Name != nil {
		// If name is empty string, reject
		if *updatedUser.Name == "" {
			writeJSONError(w, "Name cannot be empty", http.StatusBadRequest)
			return
		}
		// If name is valid, update the user (*updatedUser.Name dereferences the pointer)
		user.Name = *updatedUser.Name
	}

	users[id] = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func main() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Post("/", createUser)
		r.Get("/", getAllUsers)
		r.Delete("/{id}", deleteUserByIDHandler)
		r.Get("/{id}", getUserByIDHandler)
		r.Patch("/{id}", updateUserHandler)
	})

	http.ListenAndServe(":8080", r)

}
