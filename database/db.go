package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

const dbName = "TODO"

// GetCollection returns the MongoDB collection for the given name.
func GetCollection(name string) *mongo.Collection {
	return client.Database(dbName).Collection(name)
}

// StartMongoDB initializes the MongoDB connection.
func StartMongoDB() error {
	dbURI := os.Getenv("MONGODB_URI")
	if dbURI == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	clientOptions := options.Client().ApplyURI(dbURI)

	var err error

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	log.Println("Successfully connected to MongoDB .....")

	return nil

	// defer client.Disconnect(context.Background())
}

// CloseMongoDB closes the MongoDB connection.
func CloseMongoDB() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
