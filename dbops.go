package main

import (
	_ "database/sql"
)

// User represents a user in the database
type User struct {
	ID            int
	Name          string
	ImportantInfo string
	CreatedAt     string
}

// Entry represents an entry in the database
type Entry struct {
	ID        int
	UserID    int
	Timestamp string
	TextEntry string
}

// CreateUser inserts a new user into the database
func CreateUser(name, importantInfo string) error {
	_, err := db.Exec("INSERT INTO users (name, important_information) VALUES (?, ?)", name, importantInfo)
	return err
}

// GetUserByID retrieves a user from the database by ID
func GetUserByID(userID int) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, important_information, created_at FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Name, &user.ImportantInfo, &user.CreatedAt)
	return user, err
}

// CreateEntry inserts a new entry into the database with the current timestamp
func CreateEntry(userID int, textEntry string) error {
	_, err := db.Exec("INSERT INTO entries (user_id, timestamp, text_entry) VALUES (?, datetime('now'), ?)", userID, textEntry)
	return err
}

// GetEntriesByUserID retrieves entries for a user from the database
func GetLastNEntriesByUserID(userID, n int) ([]Entry, error) {
	rows, err := db.Query("SELECT id, user_id, timestamp, text_entry FROM entries WHERE user_id = ? order by timestamp desc limit ?", userID, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var entry Entry
		err := rows.Scan(&entry.ID, &entry.UserID, &entry.Timestamp, &entry.TextEntry)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	return entries, nil
}

// GetUserByName retrieves a user from the database by name
func GetUserByName(name string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, name, important_information, created_at FROM users WHERE name = ?", name).Scan(&user.ID, &user.Name, &user.ImportantInfo, &user.CreatedAt)
	return user, err
}
