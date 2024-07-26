package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() (string, string) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("EXERCISE_DB_API_KEY")
	connectionString := os.Getenv("GYM_CLUSTER_CONNECTION_STRING")
	return apiKey, connectionString
}
