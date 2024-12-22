package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv Loads the environment variables from the .env file
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

// Get the secret key from the environment variables
func getSecretKey() string {
	return os.Getenv("JWT_SECRET")
}
