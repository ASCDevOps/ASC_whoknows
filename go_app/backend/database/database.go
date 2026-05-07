package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"whoknows_backend/security"

	_ "github.com/jackc/pgx/v5/stdlib" // <-- pgx som database/sql driver
)

func InitDB() (*sql.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL not set")
	}

	// "pgx" er driver-navnet fra pgx/v5/stdlib
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id        SERIAL PRIMARY KEY,
        username  TEXT NOT NULL UNIQUE,
        email     TEXT NOT NULL UNIQUE,
        password  TEXT NOT NULL,
        must_change_password INTEGER NOT NULL DEFAULT 0
    );

    CREATE TABLE IF NOT EXISTS pages (
        title        TEXT PRIMARY KEY,
        url          TEXT NOT NULL UNIQUE,
        language     TEXT NOT NULL DEFAULT 'en' CHECK(language IN ('en', 'da')),
        last_updated TIMESTAMP,
        content      TEXT NOT NULL
    );
    `

	if _, err = db.Exec(schema); err != nil {
		return nil, err
	}

	if err := createAdminIfNil(db); err != nil {
		log.Printf("admin creation failed: %v\n", err)
	}

	fmt.Println("PostgreSQL connected!")
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

	var exists bool
	err := db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)",
		adminUsername,
	).Scan(&exists)
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

	_, err = db.Exec(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3)",
		adminUsername, adminEmail, hashedPassword,
	)
	if err != nil {
		return err
	}

	log.Println("Admin user created!")
	return nil
}
