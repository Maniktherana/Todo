package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Checks if env variable exists and returns env variable
func EnvMongoURI () string{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGOURI")
}