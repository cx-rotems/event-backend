package database

import (
	"database/sql"
	"log"
	"os"
	"github.com/joho/godotenv"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

const databaseUrl = "/root/database/events.db";

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = databaseUrl
	}

	DB, err = sql.Open("sqlite", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Create table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS names (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		firstName TEXT NOT NULL,
		lastName TEXT NOT NULL,
		arrived BOOLEAN DEFAULT FALSE
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
}