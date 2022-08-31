package main

import (
	"server/configs"
	"server/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize fiber app
	app := fiber.New()

	// run database
	configs.ConnectDB()

	// routes
	routes.TodoRoute(app)

	err := app.Listen(":8080")
	if err != nil {
		return
	}
}

// HI
