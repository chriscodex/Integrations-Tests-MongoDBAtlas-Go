package database_test

import (
	"testing"

	database "github.com/ChrisCodeX/CRUD-MongoDBAtlas-Go/database"
)

func TestConnection(t *testing.T) {
	_, err := database.Connection()
	if err != nil {
		t.Errorf("Database connection failed")
	}
}
