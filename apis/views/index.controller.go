package views

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-todo/configuration"
	. "go-fiber-todo/database"
	. "go-fiber-todo/database/models"
	"go-fiber-todo/utilities"
)

// Send an index view to the frontend
func IndexView(ctx *fiber.Ctx) error {
	// load records
	query := bson.D{{}}
	cursor, queryError := Instance.Database.Collection("Todos").Find(ctx.Context(), query)
	if queryError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	var todos []Todo = make([]Todo, 0)

	// iterate the cursor and decode each item into a Todo
	if err := cursor.All(ctx.Context(), &todos); err != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	// send response
	return ctx.Render("index", fiber.Map{
		"Todos": todos,
	}, "layouts/main")
}
