package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"

	"whoknows_backend/handlers"
	"whoknows_backend/database"
)

func main() {

	// Database connection
	db :=database.InitDB()
	if err != nil{
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
	mux.Handle("/api/login", &apiLoginHandler{})

	// GET /api/logout - Logout
	mux.Handle("/api/logout", &logoutHandler{})

	// Run the server on port :8080
	http.ListenAndServe(":8080", mux)
}
