package main

import (
	"go_invest/app/route"
	"go_invest/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	database.DatabaseInit()
	database.RunDatabaseMigrations()

	config := fiber.Config{
		AppName: "Go Invest",
	}
	app := fiber.New(config)
	// app.Use(cors.New())
	// app.Use(logger.New())
	// app.Use(middleware.RateLimitMiddleware)

	route.RegisterApiRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
