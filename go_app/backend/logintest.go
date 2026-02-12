package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// GET /login - Serve Login Page
	mux.Handle("/login", &loginHandler{})

	// POST /api/login - Login
	mux.Handle("/api/login", &apiLoginHandler{})

	http.ListenAndServe(":8080", mux)
}