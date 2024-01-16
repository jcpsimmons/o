// cmd/test_dbops.go
//go:build testdbops
// +build testdbops

package main

import (
	"fmt"
	"log"
	"o/dbops"
)

func main() {
	// Initialize the database connection (if needed)
	dbops.InitDB()

	// Call the functions you want to test
	users, err := dbops.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	fmt.Println("Users:", users)
}
