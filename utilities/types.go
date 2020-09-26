package utilities

import "github.com/gofiber/fiber"

type ResponseParams struct {
	Ctx    *fiber.Ctx
	Data   interface{}
	Info   string
	Status int
}
