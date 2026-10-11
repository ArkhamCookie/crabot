package database_test

import (
	"testing"

	"internal/database"
)

func TestGetDatabaseConnection(t *testing.T) {
	inputPath := "crabot.db"

	database.TestConnection(inputPath)
}
