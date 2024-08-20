package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConnectionString() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("GYM_CLUSTER_CONNECTION_STRING")
	return connectionString
}

func LoadAPIKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("EXERCISE_DB_API_KEY")
	return apiKey
}
