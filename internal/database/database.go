package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/glebarez/go-sqlite"
)

// TODO: Handle errors better

// Test connection to given database path
func TestConnection(path string) {
	if path == "" {
		path = "crabot.db"
	}

	database, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	fmt.Println("Sqlite database connection successful")
}
