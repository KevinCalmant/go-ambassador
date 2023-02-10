package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.Migrate()

	app := fiber.New()
	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		panic("Could not start server")
	}
}
