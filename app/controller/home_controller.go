package controller

import "github.com/gofiber/fiber/v2"

func ShowWelcomePageController(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Welcome to Go Invest!",
	})
}
