package configuration

import (
	"os"
)

// Database connection string
var DatabaseConnection string = os.Getenv("DATABASE_CONNECTION")

// Database name
var DatabaseName string = os.Getenv("DATABASE_NAME")

// Application port
var Port string = os.Getenv("PORT")

// Server response messages
var ResponseMessages = ResponseMessagesStruct{
	InternalServerError: "INTERNAL_SERVER_ERROR",
	MissingData:         "MISSING_DATA",
	Ok:                  "OK",
	TodoNotFound:        "TODO_NOT_FOUND",
}
