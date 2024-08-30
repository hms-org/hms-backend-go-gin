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

	APP_MODE := os.Getenv("APP_MODE")
	port := os.Getenv("APP_PORT_DEV")

	if APP_MODE == "release" {
		port = os.Getenv("APP_PORT_PROD")
	} else if APP_MODE == "testing" {
		port = os.Getenv("APP_PORT_TESTING")
	}

	return port
}
