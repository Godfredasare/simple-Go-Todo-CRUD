package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Godasare/go-todo/database"
	"github.com/Godasare/go-todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)

func FindAllTodo() []primitive.M {
	collection := database.GetCollection("todos")
	filter := bson.M{}

	var err error
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var todos []primitive.M
	for cursor.Next(context.Background()) {
		var todo bson.M
		err = cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}
		todos = append(todos, todo)
	}
	return todos
}

func InsertTodo(todo models.ToDo) {
	collection := database.GetCollection("todos")

	result, err := collection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Fatalf("Error inserting todo: %v", err)
	}

	fmt.Println("Inserted 1 movie into db:", result.InsertedID)

}

func FindOne(id string) primitive.M {
	collection := database.GetCollection("todos")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	var todo primitive.M

	filter := bson.D{{Key: "_id", Value: objID}}
	err = collection.FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		log.Fatalf("Error retrieving todo: %v", err)
	}

	fmt.Printf("found document %v", todo)
	return todo
}

func DeleteOne(id string) int64 {
	collection := database.GetCollection("todos")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatalf("Error deleting todo: %v", err)
	}

	return result.DeletedCount
}

func UpdateOne(id string, task, description *string) {
	collection := database.GetCollection("todos")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "task", Value: task},
			{Key: "description", Value: description},
		}},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(result.ModifiedCount)
}
