package main

import (
	setupport "hms-backend-go-gin/src/infra/setup_port"
	"hms-backend-go-gin/src/routes"
	"log"
)

func main() {
	app := routes.SetupRouter()
	port := setupport.SetupPort()

	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
