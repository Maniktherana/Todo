package routes

import (
	"server/controllers"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(app *fiber.App) {
	// All routes come here

	app.Post("/todo", controllers.CreateTodo)
}