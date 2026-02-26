package database

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

	

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Opens whoknows.db if null creates whoknows.db
	db, err := sql.Open("sqlite", "whoknows.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	schema := `	
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		);

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

	Println("ADMIN_USERNAME:", os.Getenv("ADMIN_USERNAME"))
	Println("ADMIN_EMAIL:", os.Getenv("ADMIN_EMAIL"))
	Println("ADMIN_PASSWORD:", os.Getenv("ADMIN_PASSWORD"))

	createAdminIfNil(db)

	// Print so we know if database is connected
	fmt.Println("SQLite connected!")

	func createAdminIfNil(db *sql.DB) {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUsername == "" || adminEmail == "" || adminPassword == "" {
		log.Println("Admin .env not set!")
		return
	}

	// Check if admin user exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", adminUsername).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		log.Println("Admin user already exists.")
		return
	}

	// Insert admin
	_, err = db.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		adminUsername,
		adminEmail,
		adminPassword,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Admin user created!")
}
