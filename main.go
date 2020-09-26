package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"

	"go-fiber-todo/configuration"
	"go-fiber-todo/utilities"
)

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
	utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Data:   todos,
		Info:   "OK",
		Status: fiber.StatusOK,
	})
}

func GetSingle(ctx *fiber.Ctx) {
	idString := ctx.Params("id")
	if idString == "" {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   "MISSING_DATA",
			Status: fiber.StatusBadRequest,
		})
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
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   "TODO_NOT_FOUND",
			Status: fiber.StatusNotFound,
		})
		return
	}

	utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Data:   element,
		Info:   "OK",
		Status: fiber.StatusOK,
	})
	return
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
	utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   "OK",
		Status: fiber.StatusOK,
	})
	return
}

func main() {
	// load environment variables
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal(envError)
		return
	}

	app := fiber.New()

	app.Use(middleware.Logger())

	app.Get("/", HandleIndex)
	app.Post("/add", AddTodo)
	app.Get("/all", GetAll)
	app.Get("/single/:id", GetSingle)
	app.Patch("/single/:id", UpdateSingle)

	// get the port
	port, portError := strconv.Atoi(configuration.Port)
	if portError != nil {
		port = 5511
	}

	// launch the app
	launchError := app.Listen(port)
	if launchError != nil {
		panic(launchError)
	}
}
