package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./database/sqlite/events.db")
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