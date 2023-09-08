package database

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const CHECK_DB_CONNECTION string = `-- name: CHECK_DB_CONNECTION :exec
SELECT 1
`

func TestDB(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error: database_test.go: godotenv.Load(): cannot load ".env" file`)
	}
	dbURL := os.Getenv("CONN")
	dbQueries, db := InitDB(dbURL)
	defer db.Close()

	t.Run("TestDB", func(t *testing.T) {
		_, err := dbQueries.db.ExecContext(context.Background(), CHECK_DB_CONNECTION)
		if err != nil {
			t.Errorf("Error: TestDB: q.ExecContext: Unable to verify connection to DB: %v", err)
			return
		}
	})

	dbQueries.TestCreateUser(t)
	dbQueries.TestGetUserByAPIKey(t)

	dbQueries.TestCreateFeed(t)
	dbQueries.TestGetAllFeeds(t)
	dbQueries.TestMarkFeedFetched(t)

	dbQueries.TestCreateFollow(t)
	dbQueries.TestDeleteFollow(t)
	dbQueries.TestGetFollowsByUser(t)

	dbQueries.TestCreatePost(t)
	dbQueries.TestGetPostsByUser(t)
}
