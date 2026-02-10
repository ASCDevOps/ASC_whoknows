package main

import (
	"html/template" // templating-pakke i go
	"net/http"      // http-pakke i go
)

var testTemplate = template.Must(template.ParseFiles("templates/test.html"))

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &homeHandler{})

	// Run the server on port :8080
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Render template
	testTemplate.Execute(w, nil)
}
