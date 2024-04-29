package controller

import (
	"go_invest/app/model"
	"go_invest/app/request"
	"go_invest/app/service"

	"github.com/gofiber/fiber/v2"
)

var userService = &service.UserService{}

func GetListUser(c *fiber.Ctx) error {
	users, err := userService.GetAllUsers()
	if err != nil {
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

func CreateUser(ctx *fiber.Ctx) error {
	req := new(request.CreateUserRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse request body",
		})
	}

	user := &model.User{
		Name:     req.Name,
		Password: req.Password,
	}

	err := userService.CreateUser(user)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data":    user,
	})
}
