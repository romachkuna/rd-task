package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"remotedevteam-task/models"
	"remotedevteam-task/routes"
)

func main() {
	startServer(":8080")
}

func startServer(listenAddress string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = models.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	// add routes to the app
	routes.ConversionRouter(app)
	err = app.Listen(listenAddress)
	if err != nil {
		log.Fatal(err)
	}
}
