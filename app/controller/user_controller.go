package controller

import (
	"github.com/gofiber/fiber/v2"
	"go_invest/app/model"
	"go_invest/database"
)

func GetListUser(c *fiber.Ctx) error {
	// get last 10 users
	var users []model.User
	result := database.DB.Select("name", "password").Limit(10).Find(&users)
	if result.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": "Failed to get data",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Hello World",
		"data":    users,
	})
}
