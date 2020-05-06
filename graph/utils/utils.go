package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// GetEnvVars - get env variables from .env
func GetEnvVars(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
