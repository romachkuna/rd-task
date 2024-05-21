package handlers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Convert() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("Romani")
	}
}
