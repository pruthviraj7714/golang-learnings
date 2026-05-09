package main

import (
	"encoding/json"
	"jwt-auth/auth"
	"jwt-auth/models"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var users []models.User

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}

	hashed, hashErr := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if hashErr != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashed)

	users = append(users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User successfully Registerd",
	})

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req models.User

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Body", http.StatusBadRequest)
		return
	}
	var usr models.User

	for _, user := range users {
		if user.Email == req.Email {
			usr = user
		}
	}

	if usr == (models.User{}) {
		http.Error(w, "User not Found", http.StatusBadRequest)
		return
	}

	er := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(req.Password))

	if er != nil {
		http.Error(w, "Invalid Password", http.StatusBadRequest)
		return
	}

	token, jwtErr := auth.GenerateJWT(usr)
	if jwtErr != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Profile Route Access Granted",
	})

}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			http.Error(w, "missing token", 401)
			return
		}

		tokenstring := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
			return auth.JWT_SECRET_KEY, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", 401)
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		if claims["role"] != "admin" {
			http.Error(w, "NotAuthorized", 403)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", registerHandler)
		r.Post("/login", loginHandler)

	})

	r.Route("/api", func(r chi.Router) {
		r.Use(authMiddleware)
		r.Get("/profile", profileHandler)

	})

	http.ListenAndServe(":8080", r)
}
