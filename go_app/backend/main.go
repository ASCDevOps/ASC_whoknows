package main

import (
	"log"
	"net/http"

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

	// Change Password
	mux.Handle("/api/change-password", &handlers.ChangePasswordHandler{DB: db})

	// GET / - Serve Root Page
	mux.Handle("/", &handlers.RootHandler{})

	// GET /register - Serve Register Page
	mux.Handle("/register", &handlers.RegisterHandler{})

	// GET /login - Serve Login Page
	mux.Handle("/login", &handlers.LoginHandler{})

	// GET /weather - Serve Weather Page
	mux.Handle("/weather", &handlers.WeatherHandler{})

	// GET /about - Serve about page
	mux.Handle("/about", &handlers.AboutHandler{})

	// GET /api/weather - Weather
	mux.Handle("/api/weather", &handlers.WeatherAPIHandler{})

	// GET /api/search - Search
	mux.Handle("/api/search", &apiSearchHandler{DB: db})

	// POST /api/register - Register
	mux.Handle("/api/register", &registerHandlerAPI{db: db})

	// POST /api/login - Login
	mux.Handle("/api/login", &handlers.APILoginHandler{DB: db})

	// GET /api/logout - Logout
	mux.Handle("/api/logout", &handlers.LogoutHandler{})

	// Serve static
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Run the server on port :8080
	http.ListenAndServe(":8080", mux)
}
