package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("EXERCISE_DB_API_KEY")
	return apiKey
}
