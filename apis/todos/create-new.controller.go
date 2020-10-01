package todos

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-todo/configuration"
	. "go-fiber-todo/database"
	. "go-fiber-todo/database/models"
	"go-fiber-todo/utilities"
)

// Create a new Todo record
func CreateNew(ctx *fiber.Ctx) error {
	// check data
	var body CreateTodoRequest
	parsingError := ctx.BodyParser(&body)
	if parsingError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	if body.Text == "" {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
	}

	collection := Instance.Database.Collection("Todos")

	// create a new record
	todo := new(Todo)
	if errorParsing := ctx.BodyParser(todo); errorParsing != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	// insert the record
	todo.Completed = false
	todo.ID = ""
	insertionResult, insertError := collection.InsertOne(ctx.Context(), todo)
	if insertError != nil {
		return utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx.Context(), filter)
	createdTodo := &Todo{}
	createdRecord.Decode(createdTodo)

	return utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Data:   createdTodo,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
}
