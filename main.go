package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

type Todo struct {
	Completed bool   `json:"completed"`
	Id        string `json:"id"`
	Text      string `json:"text"`
}

var todos = []Todo{
	{
		Completed: false,
		Id:        "1",
		Text:      "Todo number 1",
	},
	{
		Completed: true,
		Id:        "2",
		Text:      "Todo number 2",
	},
}

func GetAll(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(todos)
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("Running")
	})

	app.Get("/all", GetAll)

	error := app.Listen(3000)
	if error != nil {
		panic(error)
	}
}
