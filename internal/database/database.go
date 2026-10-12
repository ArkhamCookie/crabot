package database

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

// Test connection to given database path
func TestConnection(path string) error {
	if path == "" {
		path = "crabot.db"
	}

	database, err := sql.Open("sqlite", path)
	if err != nil {
		return err
	}

	defer database.Close()

	fmt.Println("Sqlite database connection successful")

	return nil
}

// Get SQLite version
func GetSQLiteVersion(path string) error {
	if path == "" {
		path = "crabot.db"
	}

	database, err := sql.Open("sqlite", path)
	if err != nil {
		return err
	}

	var version string
	err = database.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		return err
	}

	defer database.Close()

	fmt.Printf("SQLite Version: %s\n", version)

	return nil
}
