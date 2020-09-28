package todos

import (
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-todo/configuration"
	. "go-fiber-todo/database"
	. "go-fiber-todo/database/models"
	"go-fiber-todo/utilities"
)

// Create a new Todo record
func CreateNew(ctx *fiber.Ctx) {
	// check data
	var body CreateTodoRequest
	parsingError := ctx.BodyParser(&body)
	if parsingError != nil {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
		return
	}

	if body.Text == "" {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.MissingData,
			Status: fiber.StatusBadRequest,
		})
		return
	}

	collection := Instance.Database.Collection("Todos")

	// create a new record
	todo := new(Todo)
	if errorParsing := ctx.BodyParser(todo); errorParsing != nil {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
		return
	}

	// insert the record
	todo.Completed = false
	todo.ID = ""
	insertionResult, insertError := collection.InsertOne(ctx.Context(), todo)
	if insertError != nil {
		utilities.Response(utilities.ResponseParams{
			Ctx:    ctx,
			Info:   configuration.ResponseMessages.InternalServerError,
			Status: fiber.StatusInternalServerError,
		})
		return
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx.Context(), filter)
	createdTodo := &Todo{}
	createdRecord.Decode(createdTodo)

	utilities.Response(utilities.ResponseParams{
		Ctx:    ctx,
		Data:   createdTodo,
		Info:   configuration.ResponseMessages.Ok,
		Status: fiber.StatusOK,
	})
	return
}
