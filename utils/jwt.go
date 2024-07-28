package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func InitEnv() {
	// Load the environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func getSecretKey() string {
	// Get the secret key from the environment variables
	return os.Getenv("JWT_SECRET")
}

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // token expires after 2 hours
	})

	return token.SignedString([]byte(getSecretKey()))
}

func VerfiyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(getSecretKey()), nil
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

	userID, ok := claims["userID"].(float64) // JWT typically encodes numbers as float64

	if !ok {
		return 0, errors.New("userID is missing or not a valid number")
	}

	return int64(userID), nil
}
