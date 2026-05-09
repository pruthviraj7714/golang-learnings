package auth

import (
	"jwt-auth/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET_KEY = []byte("super-secret-key")

func GenerateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JWT_SECRET_KEY)
}
