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

// Update a single Todo record
func UpdateSingle(ctx *fiber.Ctx) error {
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
			Info:   configuration.ResponseMessages.TodoNotFound,
			Status: fiber.StatusNotFound,
		})
	}

	// parse body
	var body UpdateTodoRequest
	parsingError := ctx.BodyParser(&body)
	// TODO: parsing error can be caused by invalid data format
	if parsingError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InvalidData,
			Status: fiber.StatusBadRequest,
		})
	}

	// get collection
	collection := Instance.Database.Collection("Todos")

	// check if the record is there
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

	// update the record
	updateText := body.Text
	if updateText == "" {
		updateText = record.Text
	}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "completed", Value: body.Completed},
				{Key: "text", Value: updateText},
			},
		},
	}
	_, updateError := collection.UpdateOne(ctx.Context(), query, update)
	if updateError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
}
