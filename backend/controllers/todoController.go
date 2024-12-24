package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"backend/models"

	"github.com/gorilla/mux"
)

var todos = []models.Todo{
	{ID: "1", Title: "Learn Go", Completed: false},
	{ID: "2", Title: "Build a Todo App", Completed: false},
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = string(rand.Intn(1000000))
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedTodo models.Todo
	_ = json.NewDecoder(r.Body).Decode(&updatedTodo)
	for i, todo := range todos {
		if todo.ID == params["id"] {
			todos[i] = updatedTodo
			todos[i].ID = params["id"]
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, todo := range todos {
		if todo.ID == params["id"] {
			todos = append(todos[:i], todos[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todos)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
