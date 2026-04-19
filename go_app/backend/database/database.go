package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"whoknows_backend/security"

	// Import sqlite
	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "whoknows.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		must_change_password INTEGER NOT NULL DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS pages (
		title TEXT PRIMARY KEY UNIQUE,
		url TEXT NOT NULL UNIQUE,
		language TEXT NOT NULL CHECK(language IN ('en', 'da')) DEFAULT 'en',
		last_updated TIMESTAMP,
		content TEXT NOT NULL
	);

	-- FTS5 virtual table til hurtigere søgning i title og content
	CREATE VIRTUAL TABLE IF NOT EXISTS pages_fts USING fts5(
		title,
		content,
		language,
		content='pages',
		content_rowid='rowid'
	);

	-- Holder pages_fts opdateret når pages ændrer sig
	CREATE TRIGGER IF NOT EXISTS pages_ai AFTER INSERT ON pages BEGIN
		INSERT INTO pages_fts(rowid, title, content, language)
		VALUES (new.rowid, new.title, new.content, new.language);
	END;

	CREATE TRIGGER IF NOT EXISTS pages_ad AFTER DELETE ON pages BEGIN
		INSERT INTO pages_fts(pages_fts, rowid, title, content, language)
		VALUES('delete', old.rowid, old.title, old.content, old.language);
	END;

	CREATE TRIGGER IF NOT EXISTS pages_au AFTER UPDATE ON pages BEGIN
		INSERT INTO pages_fts(pages_fts, rowid, title, content, language)
		VALUES('delete', old.rowid, old.title, old.content, old.language);

		INSERT INTO pages_fts(rowid, title, content, language)
		VALUES (new.rowid, new.title, new.content, new.language);
	END;
	`

	_, err = db.Exec(schema)
	if err != nil {
		return nil, err
	}

	// Fylder pages_fts med eksisterende data fra pages
	_, err = db.Exec(`INSERT INTO pages_fts(pages_fts) VALUES('rebuild');`)
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

	hashedPassword, err := security.HashPassword(adminPassword)
	if err != nil {
		return fmt.Errorf("failed to hash admin password: %w", err)
	}

	// Insert admin
	_, err = db.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		adminUsername,
		adminEmail,
		hashedPassword,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("Admin user created!")
	return nil
}
