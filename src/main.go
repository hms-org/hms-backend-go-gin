package main

import (
	setupport "hms-backend-go-gin/infra/setup_port"
	"hms-backend-go-gin/routes"
)

func main() {
	app := routes.SetupRouter()
	port := setupport.SetupPort()

	app.Run(":" + port)
}
