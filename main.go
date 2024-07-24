package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id int
	Message string
}

func main() {
	data := map[string][]Todo{
		"Todos": {
			{Id: 1, Message: "Buy milk"},
		},
	}
	

// todosHandler is a handler function that renders the index.html template
	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		templ.Execute(w, data)
	}

	
	// addTodoHandler is a handler function that adds a new todo to the list
	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		templ := template.Must(template.ParseFiles("index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	// The code below starts a web server that listens on port 8000. It has two routes: one for rendering the index.html template and another for adding a new todo to the list.
	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}