package main

import (
	"o/dbops"
)

func main() {
	dbops.InitDB()
	// slice of names
	names := []string{"Walter White", "Jessie Pinkman", "Frankie D. Dog."}
	// for each name create db user
	for _, name := range names {
		// create user
		err := dbops.CreateUser(name, "This is important information")
		if err != nil {
			panic(err)
		}
	}
}
