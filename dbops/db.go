package dbops

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	database, err := sql.Open("sqlite3", "./userdata.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}

	db = database

	createTables()
	fmt.Println("Database connection established.")
}

func createTables() {
	// create users table and entries
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			important_information TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS entries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			text_entry TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
    `

	_, err := db.Exec(query)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Table created successfully.")
}

// Exec is a placeholder for the database execution method
func Exec(query string, args ...interface{}) (sql.Result, error) {
	// Your database execution code...
	return nil, nil
}

func Close() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed.")
	}
}
