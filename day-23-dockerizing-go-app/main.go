package main

import (
	"dockerizing-go-app/db"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Healthy Server",
	})
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Error loading .env file: %s", err.Error())
	}

	r := chi.NewRouter()

	database := db.Connect()

	// Migrate the schema with GORM
	database.AutoMigrate(&User{})

	fmt.Println("Database migrated")

	r.Get("/", rootHandler)
	// Add User to database
	r.Post("/add-user", func(w http.ResponseWriter, r *http.Request) {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			fmt.Printf("Error decoding JSON: %s", err.Error())
			return
		}

		database.Create(&user)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]User{
			"user": user,
		})
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		var users []User

		database.Find(&users)

		if len(users) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "No users found",
			})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string][]User{
			"users": users,
		})
	})

	fmt.Println("Server is running on 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

}
