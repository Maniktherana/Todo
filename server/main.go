package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize fiber app
	app := fiber.New()

	// Handler func returns JSON 
	app.Get("/", func(c *fiber.Ctx) error {
		// &fiber.Map  = map[string]interface{} for JSON returns
		return c.JSON(&fiber.Map{"data": "Hello from fiber & mongodb"})
	})

	app.Listen(":6000")
}