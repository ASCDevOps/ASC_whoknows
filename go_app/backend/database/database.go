package database

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

	
func InitDB() (*sql.DB, error) {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Opens whoknows.db if null creates whoknows.db
	db, err := sql.Open("sqlite", "whoknows.db")
	if err != nil {
		return nil, err
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

	createAdminIfNil(db)

	// Print so we know if database is connected
	fmt.Println("SQLite connected!")

return db, nil

}

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
		log.Println(err)
		return
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

	if err != nil{
		log.Println(err)
		return
	}

	log.Println("Admin user created!")
}