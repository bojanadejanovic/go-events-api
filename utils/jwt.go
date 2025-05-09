package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int64, email string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userID": userID,
			"exp":    time.Now().Add(time.Hour * 2).Unix(),
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
	if secret != "" {
		fmt.Println("secret retrieved")
	}
	return token.SignedString([]byte(secret))
}

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	// Convert float64 to int64 for userID
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("invalid userID in token")
	}
	userID := int64(userIDFloat)

	return userID, nil
}
