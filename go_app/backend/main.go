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
	db, err := sql.Open("sqlite", "whoknows.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := `
	DROP TABLE IF EXISTS users;

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);

	INSERT INTO users (username, email, password) 
	VALUES ('admin', 'keamonk1@stud.kea.dk', '5f4dcc3b5aa765d61d8327deb882cf99');

	CREATE TABLE IF NOT EXISTS pages (
		title TEXT PRIMARY KEY UNIQUE,
		url TEXT NOT NULL UNIQUE,
		language TEXT NOT NULL CHECK(language IN ('en', 'da')) DEFAULT 'en',
		last_updated TIMESTAMP,
		content TEXT NOT NULL
	);`

	_, err = db.Exec(schema)
	if err != nil {
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
