package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go_invest/app/route"
	"go_invest/database"
)

func main() {
	database.DatabaseInit()
	database.RunDatabaseMigrations()

	app := fiber.New()

	route.RegisterApiRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
