package utils

import (
	"gin-app/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWTToken(storedUser models.User) (string, error) {
	// Define the JWT claims
	claims := jwt.MapClaims{
		"user_id": storedUser.ID,
		"email":   storedUser.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := os.Getenv("JWT")
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
