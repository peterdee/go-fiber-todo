package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go-fiber-todo/configuration"
)

var mg MongoInstance

// const mongoURI = "mongodb://user:password@localhost:27017/" + dbName

// Connect configures the MongoDB client and initializes the database connection.
// Source: https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup
func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(configuration.DatabaseConnection))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(configuration.DatabaseName)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client:   client,
		Database: db,
	}

	return nil
}
