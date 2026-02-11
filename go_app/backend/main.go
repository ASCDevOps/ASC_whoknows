package main

import (
	"encoding/json" // Needed for endpoints
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
	// TODO: Implement ^

	// GET /api/search - Search
	// TODO: Implement ^

	// POST /api/register - Register
	// TODO: Implement ^

	// POST /api/login - Login
	// TODO: Implement ^

	// GET /api/logout - Logout
	mux.Handle("/api/logout", &logoutHandler{})

	// Run the server on port :8080
	http.ListenAndServe(":8080", mux)
}

type rootHandler struct{}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve Root Page
	testTemplate.Execute(w, nil)
}

type logoutHandler struct{}

func (h *logoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// JSON-response
	resp := map[string]string{
		"status": "logged out",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
