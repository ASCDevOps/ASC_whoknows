package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http" // http-pakke in go

	_ "modernc.org/sqlite"
)

func main() {

	// Create a new request multiplexer

	//opens whoknows.db if null creates whoknows.db
	db, err := sql.Open("sqlite", "file:whoknows.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//error handling
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	//print so we know if database is connected
	fmt.Println("SQLite connected!")
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// GET / - Serve Root Page
	mux.Handle("/", &rootHandler{})

	// GET /register - Serve Register Page
	mux.Handle("/register", &registerHandler{})

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
