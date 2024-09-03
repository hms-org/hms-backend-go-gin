package setupport

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("APP_PORT")
	return port
}
