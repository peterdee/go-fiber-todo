package todos

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-fiber-todo/configuration"
	. "go-fiber-todo/database"
	"go-fiber-todo/utilities"
)

// Update a single Todo record
func DeleteSingle(ctx *fiber.Ctx) error {
	// check data
	id := ctx.Params("id")
	if id == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}

	// parse Todo ID
	todoId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InvalidData,
			Status: fiber.StatusNotFound,
		})
	}

	// get collection
	collection := Instance.Database.Collection("Todos")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: todoId}}
	result, deleteError := collection.DeleteOne(ctx.Context(), &query)
	if deleteError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	// check if item was deleted
	if result.DeletedCount < 1 {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.TodoNotFound,
			Status: fiber.StatusNotFound,
		})
	}

	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
}
