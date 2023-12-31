package app

import (
	"log"

	"github.com/joho/godotenv"
)

// This function loads the `.env` file using `joho/godotenv`
func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
