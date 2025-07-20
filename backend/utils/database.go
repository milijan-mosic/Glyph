package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (skipping)")
	}

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}

	return value
}

func GetDatabaseUrl() string {
	username := GetEnvVariable("POSTGRES_USER")
	password := GetEnvVariable("POSTGRES_PASSWORD")
	dbName := GetEnvVariable("POSTGRES_DB")

	url := fmt.Sprintf("postgres://%s:%s@database:5432/%s", username, password, dbName)

	return url
}
