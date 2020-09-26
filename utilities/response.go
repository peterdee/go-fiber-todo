package utilities

import "github.com/gofiber/fiber"

func Response(params ResponseParams) {
	// TODO: don't send the 'data' field if there's no data

	data := params.Data
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
			"data":   data,
			"info":   info,
			"status": status,
		},
	)
	return
}
