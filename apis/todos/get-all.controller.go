package todos

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-todo/configuration"
	. "go-fiber-todo/database"
	. "go-fiber-todo/database/models"
	"go-fiber-todo/utilities"
)

// get all Todo records
func GetAll(ctx *fiber.Ctx) error {
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

	// iterate the cursor and decode each item into an Employee
	if err := cursor.All(ctx.Context(), &todos); err != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	// send response
	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Data:   todos,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
}
