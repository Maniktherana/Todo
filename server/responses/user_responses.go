package responses

import (
	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	Status	int			`json:"status"`
	Message	string		`json:"message"`
	Data	*fiber.Map	`json:"data"`
}
// The snippet above creates a UserResponse struct 
// with Status, Message, and Data property to represent 
// the API response type.