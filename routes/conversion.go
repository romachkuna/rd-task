package routes

import (
	"github.com/gofiber/fiber/v2"
	"remotedevteam-task/handlers"
)

func ConversionRouter(app fiber.Router) {
	app.Post("/convert", handlers.Convert())
}
