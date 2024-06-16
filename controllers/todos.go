package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Godasare/go-todo/models"
	"github.com/Godasare/go-todo/services"
)

func HomeServer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my todo app"))
}

func GetAllTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post pruduct")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	todo := services.FindAllTodo()
	json.NewEncoder(w).Encode(todo)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post pruduct")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Alow-Methods", "POST")

	var todo models.ToDo
	todo.Created_At = time.Now().Format("2006-01-02 3:4:5 pm")
	json.NewDecoder(r.Body).Decode(&todo)
	services.InsertTodo(todo)
	json.NewEncoder(w).Encode(todo)
}

func GetOneTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get product by ID")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")

	id := r.PathValue("id")
	todo := services.FindOne(id)
	json.NewEncoder(w).Encode(todo)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete product by ID")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	id := r.PathValue("id")
	deleteCount := services.DeleteOne(id)
	if deleteCount == 0 {
		http.Error(w, "Do todo found with the given ID", http.StatusBadRequest)
	} else {

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Deleted Sucessfully"})
	}

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete product by ID")
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")

	var todo models.ToDo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id := r.PathValue("id")
	services.UpdateOne(id, &todo.Task, &todo.Description)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Updated Sucessfully"})

}
