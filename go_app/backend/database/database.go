package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	// Import sqlite
	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {

	// Load .env file
	// Ignoring error on purpose, for production purposes.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		// fallback
		dbPath = "whoknows.db"
	}

	// Opens whoknows.db if null creates whoknows.db
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

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
		return nil, err
	}

	if err := createAdminIfNil(db); err != nil {
		log.Printf("admin creation failed: %v\n", err)
	}

	// Print so we know if database is connected
	fmt.Println("SQLite connected!")

	return db, nil

}

func createAdminIfNil(db *sql.DB) error {
	adminUsername := os.Getenv("ADMIN_USERNAME")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminUsername == "" || adminEmail == "" || adminPassword == "" {
		log.Println("Admin .env not set!")
		return nil
	}

	// Check if admin user exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", adminUsername).Scan(&exists)

	if err != nil {
		return fmt.Errorf("failed to check for existing admin user: %w", err)
	}

	if exists {
		log.Println("Admin user already exists.")
		return nil
	}

	// Insert admin
	_, err = db.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		adminUsername,
		adminEmail,
		adminPassword,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Admin user created!")
	return nil
}
