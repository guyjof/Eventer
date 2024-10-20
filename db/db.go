package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "eventer.db")

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	DB.SetMaxOpenConns(10) // max number of open connections
	DB.SetMaxIdleConns(5)  // min number of idle connections

	// Ping the database to ensure the connection is valid
	if err := DB.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	createTables()
}

func createTables() {
	createUsersTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(createUsersTableQuery)
	if err != nil {
		panic(fmt.Sprintf("Failed to create users table: %v\n", err))
	}

	createEventsTableQuery := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,	
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createEventsTableQuery)
	if err != nil {
		panic(fmt.Sprintf("Failed to create events table: %v\n", err))
	}

	createRegistrationsTableQuery := `
    CREATE TABLE IF NOT EXISTS registrations (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      user_id INTEGER,
      event_id INTEGER,
      FOREIGN KEY(user_id) REFERENCES users(id),
      FOREIGN KEY(event_id) REFERENCES events(id)
    )
  `
	_, err = DB.Exec(createRegistrationsTableQuery)
	if err != nil {
		panic(fmt.Sprintf("Failed to create registrations table: %v\n", err))
	}
}
