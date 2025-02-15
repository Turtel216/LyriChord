package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file found on the given file. Panics and logs error on fail
func loadEnv(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}

// GetEnvVar returns the .env value for the given key and the path to the .env file.
// If the key is not found it returns an empty string and an error
func GetEnvVar(path string, key string) (string, error) {
	loadEnv(path)
	value := os.Getenv(key)

	if value == "" {
		return value, errors.New("Invalid .env key")
	}

	return value, nil
}
