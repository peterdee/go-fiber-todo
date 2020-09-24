package main

import (
	"strconv"

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

func AddTodo(ctx *fiber.Ctx) {
	type request struct {
		Text string `json:"text"`
	}

	var body request
	error := ctx.BodyParser(&body)
	if error != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": error,
			},
		)
		return
	}

	if body.Text == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"info": "MISSING_DATA",
			},
		)
		return
	}

	todo := Todo{
		Completed: false,
		Id:        strconv.Itoa(len(todos) + 1),
		Text:      body.Text,
	}

	todos = append(todos, todo)
	ctx.Status(fiber.StatusOK).JSON(todos)
}

func GetAll(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(todos)
}

func GetSingle(ctx *fiber.Ctx) {
	idString := ctx.Params("id")
	if idString == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"info": "MISSING_DATA",
			},
		)
		return
	}

	var element Todo
	for i := range todos {
		if todos[i].Id == idString {
			element = todos[i]
			break
		}
	}

	if element.Id == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"info": "TODO_NOT_FOUND",
			},
		)
		return
	}

	ctx.Status(fiber.StatusOK).JSON(element)
}

func UpdateSingle(ctx *fiber.Ctx) {
	idString := ctx.Params("id")
	if idString == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"info": "MISSING_DATA",
			},
		)
		return
	}

	type request struct {
		Completed bool   `json:"completed"`
		Text      string `json:"text"`
	}
	var body request
	error := ctx.BodyParser(&body)
	if error != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": error,
			},
		)
		return
	}

	var element Todo
	for i := range todos {
		if todos[i].Id == idString {
			todos[i].Completed = body.Completed
			todos[i].Text = body.Text
			element = todos[i]
			break
		}
	}

	if element.Id == "" {
		ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"info": "TODO_NOT_FOUND",
			},
		)
		return
	}

	ctx.Status(fiber.StatusOK).JSON(element)
}

func HandleIndex(ctx *fiber.Ctx) {
	ctx.Send("Running")
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/", HandleIndex)
	app.Post("/add", AddTodo)
	app.Get("/all", GetAll)
	app.Get("/single/:id", GetSingle)
	app.Patch("/single/:id", UpdateSingle)

	error := app.Listen(3000)
	if error != nil {
		panic(error)
	}
}
