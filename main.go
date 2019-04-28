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

var todosFile = "todos.json"
var r = mux.NewRouter()
var todos []Todo

// Home is our main endpoint to view all cruds and query CRUD on todos
func Home(w http.ResponseWriter, r *http.Request) {

}

// GetTodos return all registred todos in todosFile
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
	jsonContent, err := openTodoFile()
	params := mux.Vars(r)
	var todo Todo

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	json.Unmarshal(jsonContent, &todos)
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = params["id"]
	todos = append(todos, todo)

	saveTodoFile(todos)
	json.NewEncoder(w).Encode(todos)
}

// DeleteTodo return a single todo based on todo id
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	jsonContent, err := openTodoFile()

	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	var newTodos []Todo
	json.Unmarshal(jsonContent, &todos)

	params := mux.Vars(r)
	for _, todo := range todos {
		if todo.ID == params["id"] {
			continue
		}

		newTodos = append(newTodos, todo)
	}

	saveTodoFile(newTodos)
	json.NewEncoder(w).Encode(newTodos)
}

func registerRoutes() {
	// Main routes
	r.HandleFunc("/", Home)

	var todosRouter = r.PathPrefix("/todos").Subrouter()
	todosRouter.HandleFunc("/", GetTodos).Methods("GET")
	todosRouter.HandleFunc("/{id}", GetTodo).Methods("GET")
	todosRouter.HandleFunc("/{id}", DeleteTodo).Methods("DELETE")
	todosRouter.HandleFunc("/{id}", CreateTodo).Methods("POST")

	// Static content
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	http.Handle("/", r)

	// Headers
	r.Headers("Content-Type", "application/json")
}

// OpenTodoFile return todos.json content
func openTodoFile() ([]byte, error) {
	jsonFile, err := os.Open(todosFile)

	if err != nil {
		return make([]byte, 0), errors.New("Error opening todos file")
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue, nil
}

func saveTodoFile(newTodos []Todo) {
	newJSONContent, _ := json.Marshal(newTodos)
	ioutil.WriteFile(todosFile, newJSONContent, 0644)
}

func main() {
	registerRoutes()
	http.ListenAndServe(":8080", r)
}
