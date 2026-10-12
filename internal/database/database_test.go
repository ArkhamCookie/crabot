package database_test

import (
	"testing"

	"internal/database"
)

func TestGetDatabaseConnection(t *testing.T) {
	inputPath := "crabot.db"

	database.TestConnection(inputPath)
}

func TestGetSQLiteVersion(t *testing.T) {
	inputPath := "crabot.db"

	err := database.GetSQLiteVersion(inputPath)
	if err != nil {
		t.Fatal("Error:", err)
	}
}
