package database

import (
	//"context"
	//"database/sql"
	"log"
	"os"
	"testing"
	//"time"

	//"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func TestDB(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error: database_test.go: godotenv.Load(): cannot load ".env" file`)
	}
	dbURL := os.Getenv("CONN")
	dbQueries, db := InitDB(dbURL)
	defer db.Close()

	dbQueries.TestCreateUser(t)
	dbQueries.TestGetUserByAPIKey(t)

	dbQueries.TestCreateFeed(t)
	dbQueries.TestGetAllFeeds(t)
	dbQueries.TestMarkFeedFetched(t)
}
