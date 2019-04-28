package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Todo represents a list of created todos
type Todo struct {
	ID          string
	Description string
	Finished    bool
	Date        string
}

var r = mux.NewRouter()
var todos []Todo

// Home is our main endpoint to view all cruds and query CRUD on todos
func Home(w http.ResponseWriter, r *http.Request) {

}

// GetTodos return all registred todosin todos.json
func GetTodos(w http.ResponseWriter, r *http.Request) {
	jsonContent, err := openTodoFile()

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	json.Unmarshal(jsonContent, &todos)
	json.NewEncoder(w).Encode(todos)
}

// GetTodo return a single todo based on todo id
func GetTodo(w http.ResponseWriter, r *http.Request) {
	jsonContent, err := openTodoFile()

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	json.Unmarshal(jsonContent, &todos)
	params := mux.Vars(r)
	for _, todo := range todos {
		if todo.ID == params["id"] {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	json.NewEncoder(w).Encode(&Todo{})
}

// CreateTodo return a single todo based on todo id
func CreateTodo(w http.ResponseWriter, r *http.Request) {

}

// DeleteTodo return a single todo based on todo id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {

}

func registerRoutes() {
	// Main routes
	r.HandleFunc("/", Home)

	var todosRouter = r.PathPrefix("/todos").Subrouter()
	todosRouter.HandleFunc("/", GetTodos).Methods("GET")
	todosRouter.HandleFunc("/{id}", GetTodo).Methods("GET")
	todosRouter.HandleFunc("/{id}", DeleteTodo).Methods("DELETE")
	todosRouter.HandleFunc("/", CreateTodo).Methods("POST")

	// Static content
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	http.Handle("/", r)

	// Headers
	r.Headers("Content-Type", "application/json")
}

// OpenTodoFile return todos.json content
func openTodoFile() ([]byte, error) {
	jsonFile, err := os.Open("todos.json")

	if err != nil {
		return make([]byte, 0), errors.New("Error opening todos file")
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}

func main() {
	registerRoutes()
	http.ListenAndServe(":8080", r)
}
