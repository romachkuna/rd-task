package handlers

import (
	"github.com/gofiber/fiber/v2"
	"remotedevteam-task/conversion_api"
	"remotedevteam-task/models"
)

func Convert() fiber.Handler {
	return func(c *fiber.Ctx) error {
		baseCurrency := c.Query("base_currency")
		targetCurrency := c.Query("target_currency")

		if baseCurrency == "" || targetCurrency == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "base_currency and target_currency are required",
			})
		}
		var r models.Request
		err := c.BodyParser(&r)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input values",
			})
		}
		r.BaseCurrency = baseCurrency
		r.TargetCurrency = targetCurrency
		convertedAmount, err := conversion_api.Fetch(r)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		conversionResponse := models.ConversionResponse{
			BaseCurrency:       r.BaseCurrency,
			BaseCurrencyAmount: r.Amount,
			TargetCurrency:     r.TargetCurrency,
			ConvertedAmount:    convertedAmount,
		}
		return c.Status(fiber.StatusOK).JSON(conversionResponse)
	}
}
