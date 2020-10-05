package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"

	"go-fiber-todo/apis/index"
	"go-fiber-todo/apis/todos"
	"go-fiber-todo/apis/views"
	"go-fiber-todo/configuration"
	"go-fiber-todo/database"
	"go-fiber-todo/utilities"
)

func main() {
	// load environment variables via the .env file
	env := os.Getenv("ENV")
	if env != "heroku" {
		envError := godotenv.Load()
		if envError != nil {
			log.Fatal(envError)
			return
		}
	}

	// connect to the database
	dbError := database.Connect()
	if dbError != nil {
		log.Fatal(dbError)
		return
	}

	// create a new views engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// middlewares
	app.Use(logger.New())

	// serve static files
	app.Static("/", "./public")

	// views
	app.Get("/", views.IndexView)

	// available APIs
	app.Get("/api", index.IndexController)
	app.Post("/api/todos/add", todos.CreateNew)
	app.Get("/api/todos/all", todos.GetAll)
	app.Delete("/api/todos/delete/:id", todos.DeleteSingle)
	app.Get("/api/todos/get/:id", todos.GetSingle)
	app.Patch("/api/todos/update/:id", todos.UpdateSingle)

	// handle 404
	app.Use(func(ctx *fiber.Ctx) error {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.NotFound,
			Status: fiber.StatusNotFound,
		})
	})

	// get the port
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5511"
	}

	fmt.Println("port", port)
	// launch the app
	launchError := app.Listen(port)
	if launchError != nil {
		panic(launchError)
	}
}
