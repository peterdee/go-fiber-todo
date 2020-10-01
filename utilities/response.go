package utilities

import (
	"github.com/gofiber/fiber/v2"
)

// Send a response
func Response(params ResponseParams) error {
	info := params.Info
	status := params.Status
	if info == "" {
		info = "OK"
	}
	if status == 0 {
		status = 200
	}
	return params.Ctx.Status(params.Status).JSON(
		fiber.Map{
			"data":   params.Data,
			"info":   info,
			"status": status,
		},
	)
}
