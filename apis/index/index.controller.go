package index

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-todo/configuration"
	"go-fiber-todo/utilities"
)

// Handle index route
func IndexController(ctx *fiber.Ctx) error {
	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
}
