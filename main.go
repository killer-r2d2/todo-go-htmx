package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

var todos []Todo

func main() {
	// Initialize with a sample todo
	todos = []Todo{
		{Id: 1, Message: "Buy milk"},
	}

	// Handler to render the index.html template
	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		data := map[string]interface{}{
			"Todos": todos,
		}
		templ.Execute(w, data)
	}

	// Handler to add a new todo item
	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		todo := Todo{Id: len(todos) + 1, Message: message}
		todos = append(todos, todo)

		// Render only the new todo item, using a small template string
		tmpl, _ := template.New("todo").Parse(`<li>{{.Message}}</li>`)
		tmpl.Execute(w, todo)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
