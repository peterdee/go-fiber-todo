package utilities

import (
	"github.com/gofiber/fiber"
)

// Send a response
func Response(params ResponseParams) {
	info := params.Info
	status := params.Status
	if info == "" {
		info = "OK"
	}
	if status == 0 {
		status = 200
	}
	params.Ctx.Status(params.Status).JSON(
		fiber.Map{
			"data":   params.Data,
			"info":   info,
			"status": status,
		},
	)
	return
}
