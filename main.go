package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/joho/godotenv"

	"go-fiber-todo/apis/index"
	"go-fiber-todo/apis/todos"
	"go-fiber-todo/configuration"
	"go-fiber-todo/database"
	"go-fiber-todo/utilities"
)

var todoss = []Todo{
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
		Id:        strconv.Itoa(len(todoss) + 1),
		Text:      body.Text,
	}

	todoss = append(todoss, todo)
	ctx.Status(fiber.StatusOK).JSON(todoss)
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
	for i := range todoss {
		if todoss[i].Id == idString {
			element = todoss[i]
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
	for i := range todoss {
		if todoss[i].Id == idString {
			todoss[i].Completed = body.Completed
			todoss[i].Text = body.Text
			element = todoss[i]
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

func main() {
	// load environment variables
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal(envError)
		return
	}

	// connect to the database
	dbError := database.Connect()
	if dbError != nil {
		log.Fatal(dbError)
		return
	}

	app := fiber.New()

	// middlewares
	app.Use(middleware.Logger())

	// available APIs
	app.Get("/", index.IndexController)
	app.Post("/add", AddTodo)
	app.Get("/api/todos/all", todos.GetAll)
	app.Get("/single/:id", GetSingle)
	app.Patch("/single/:id", UpdateSingle)

	// handle 404
	app.Use(func(ctx *fiber.Ctx) {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.NotFound,
			Status: fiber.StatusNotFound,
		})
	})

	// get the port
	port := os.Getenv("PORT")
	portInt, portError := strconv.Atoi(port)
	if portError != nil {
		portInt = 5511
	}

	// launch the app
	launchError := app.Listen(portInt)
	if launchError != nil {
		panic(launchError)
	}
}
