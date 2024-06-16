package router

import (
	"net/http"

	"github.com/Godasare/go-todo/controllers"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", controllers.HomeServer)
	router.HandleFunc("POST /api/todo", controllers.CreateTodo)
	router.HandleFunc("GET /api/todo", controllers.GetAllTodo)
	router.HandleFunc("GET /api/todo/{id}", controllers.GetOneTodo)
	router.HandleFunc("DELETE /api/todo/{id}", controllers.DeleteTodo)
	router.HandleFunc("PUT /api/todo/{id}", controllers.UpdateTodo)

	return router
}
