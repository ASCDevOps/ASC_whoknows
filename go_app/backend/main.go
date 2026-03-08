package main

import (
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"

	"whoknows_backend/database"
	"whoknows_backend/handlers"
)

func main() {
	// Database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// GET / - Serve Root Page
	mux.Handle("/", &handlers.RootHandler{})

	// GET /register - Serve Register Page
	mux.Handle("/register", &handlers.RegisterHandler{})

	// GET /login - Serve Login Page
	mux.Handle("/login", &handlers.LoginHandler{})

	// GET /weather - Serve Weather Page
	mux.Handle("/weather", &handlers.WeatherHandler{})

	// GET /api/weather - Weather
	mux.Handle("/api/weather", &handlers.WeatherAPIHandler{})

	// GET /api/search - Search
	mux.Handle("/api/search", &apiSearchHandler{DB: db})

	// POST /api/register - Register
	mux.Handle("/api/register", &registerHandlerAPI{db: db})

	// POST /api/login - Login
	mux.Handle("/api/login", &handlers.APILoginHandler{DB: db})

	// GET /api/logout - Logout
	mux.Handle("/api/logout", &logoutHandler{})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}