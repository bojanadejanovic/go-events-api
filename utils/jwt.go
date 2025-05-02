package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int64, email string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"sub":   userID,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		},
	)

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Try getting JWT_SECRET again after loading .env.local
		secret = os.Getenv("JWT_SECRET")
		if secret == "" {
			return "", fmt.Errorf("JWT_SECRET environment variable not set")
		}
	}
	return token.SignedString([]byte(secret))
}
