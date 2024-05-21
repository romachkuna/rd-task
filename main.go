package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"remotedevteam-task/models"
	"remotedevteam-task/routes"
)

func main() {
	startServer(":8080")
}

func startServer(listenAddress string) {
	err := models.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routes.ConversionRouter(app)
	err = app.Listen(listenAddress)
	if err != nil {
		log.Fatal(err)
	}
}
