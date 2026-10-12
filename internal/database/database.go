package database

import (
	"database/sql"

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

	return nil
}

// Get SQLite version
func GetSQLiteVersion(path string) (string, error) {
	if path == "" {
		path = "crabot.db"
	}

	database, err := sql.Open("sqlite", path)
	if err != nil {
		return "", err
	}

	var version string
	err = database.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	defer database.Close()

	return version, nil
}
