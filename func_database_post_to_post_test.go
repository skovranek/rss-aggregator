package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/skovranek/rss_aggregator/internal/database"
)

func TestDBPostToPost(t *testing.T) {
	id := uuid.New()

	now := time.Now().UTC()
	nullTimeNow := sql.NullTime{
		Time:  now,
		Valid: true,
	}

	str := "this is a string"
	nullStr := sql.NullString{
		String: str,
		Valid:  true,
	}

	tests := []struct {
		input  database.Post
		expect Post
	}{
		{}, // test zero values
		{
			input: database.Post{
				ID:          id,
				CreatedAt:   now,
				UpdatedAt:   now,
				Title:       nullStr,
				Url:         str,
				Description: nullStr,
				PublishedAt: nullTimeNow,
				FeedID:      id,
			},
			expect: Post{
				ID:          id,
				CreatedAt:   now,
				UpdatedAt:   now,
				Title:       str,
				Url:         str,
				Description: str,
				PublishedAt: now,
				FeedID:      id,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestDBPostToPost Case #%v:", i), func(t *testing.T) {
			output := databasePostToPost(test.input)
			if output != test.expect {
				t.Errorf("Unexpected: TestDBPostToPost: \n%v", output)
				return
			}
		})
	}
}
