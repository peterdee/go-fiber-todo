package todos

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-fiber-todo/configuration"
	. "go-fiber-todo/database"
	. "go-fiber-todo/database/models"
	"go-fiber-todo/utilities"
)

// Get a single Todo record
func GetSingle(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}

	todoId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.NotFound,
			Status: fiber.StatusNotFound,
		})
	}

	collection := Instance.Database.Collection("Todos")

	query := bson.D{{Key: "_id", Value: todoId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &Todo{}
	rawRecord.Decode(record)

	if record.ID == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.NotFound,
			Status: fiber.StatusNotFound,
		})
	}

	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Data:   record,
		Info:   "OK",
		Status: fiber.StatusOK,
	})
}
