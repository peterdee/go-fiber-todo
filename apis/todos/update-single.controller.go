package todos

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go-fiber-todo/configuration"
	. "go-fiber-todo/database"
	. "go-fiber-todo/database/models"
	"go-fiber-todo/utilities"
)

// Update a single Todo record
func UpdateSingle(ctx *fiber.Ctx) {
	// check data
	id := ctx.Params("id")
	if id == "" {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
		return
	}

	// parse Todo ID
	todoId, parseError := primitive.ObjectIDFromHex(id)
	if parseError != nil {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.TodoNotFound,
			Status: fiber.StatusNotFound,
		})
		return
	}

	// parse body
	var body UpdateTodoRequest
	parsingError := ctx.BodyParser(&body)
	// TODO: parsing error can be caused by invalid data format
	if parsingError != nil {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InvalidData,
			Status: fiber.StatusBadRequest,
		})
		return
	}

	// get collection
	collection := Instance.Database.Collection("Todos")

	// check if the record is there
	query := bson.D{{Key: "_id", Value: todoId}}
	rawRecord := collection.FindOne(ctx.Context(), query)
	record := &Todo{}
	rawRecord.Decode(record)
	if record.ID == "" {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.NotFound,
			Status: fiber.StatusNotFound,
		})
		return
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
	_, updateError := collection.UpdateOne(ctx.Fasthttp, query, update)
	if updateError != nil {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
		return
	}

	utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
	return
}
