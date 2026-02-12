package main

import (
	"html/template" // templating-pakke in go
	"net/http"      // http-pakke in go
)

var testTemplate = template.Must(template.ParseFiles("templates/test.html"))

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// GET / - Serve Root Page
	mux.Handle("/", &rootHandler{})

	// GET /register - Serve Register Page
	// TODO: Implement ^

	// GET /login - Serve Login Page
	mux.Handle("/login", &loginHandler{})

	// GET /api/search - Search
	// TODO: Implement ^

	// POST /api/register - Register
	// TODO: Implement ^

	// POST /api/login - Login
	mux.Handle("/api/login", &apiLoginHandler{})

	// GET /api/logout - Logout
	mux.Handle("/api/logout", &logoutHandler{})

	// Run the server on port :8080
	http.ListenAndServe(":8080", mux)
}
