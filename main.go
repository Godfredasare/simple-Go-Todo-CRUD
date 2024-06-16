package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Godasare/go-todo/config"
	"github.com/Godasare/go-todo/database"
	"github.com/Godasare/go-todo/router"
)

func init() {
	config.LoadEnv()

	err := database.StartMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

}

func main() {
	PORT := os.Getenv("PORT")

	fmt.Println("My Todo App")

	r := router.Router()

	err := http.ListenAndServe(":"+PORT, r)
	if err != nil {
		panic(err)
	}

	defer database.CloseMongoDB()
}
