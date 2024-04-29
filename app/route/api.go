package route

import (
	"github.com/gofiber/fiber/v2"
	"go_invest/app/controller"
)

func RegisterApiRoutes(route *fiber.App) {
	route.Get("/", controller.ShowWelcomePageController)

	route.Get("/users", controller.GetListUser)
	route.Post("/users", controller.CreateUser)
}
