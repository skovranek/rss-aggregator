package main

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/skovranek/rss_aggregator/internal/database"
)

func TestDBFeedToFeed(t *testing.T) {
	id := uuid.New()

	now := time.Now().UTC()
	str := "this is a string"

	tests := []struct {
		input  database.Feed
		expect Feed
	}{
		{
			input: database.Feed{
				ID:        id,
				CreatedAt: now,
				UpdatedAt: now,
				Name:      str,
				Url:       str,
				UserID:    id,
				LastFetchedAt: sql.NullTime{
					Time: now,
				},
			},
			expect: Feed{
				ID:            id,
				CreatedAt:     now,
				UpdatedAt:     now,
				Name:          str,
				Url:           str,
				UserID:        id,
				LastFetchedAt: now,
			},
		},
		{
			input:  database.Feed{},
			expect: Feed{},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestDBFeedToFeed Case #%v:", i), func(t *testing.T) {
			output := databaseFeedToFeed(test.input)
			if output != test.expect {
				t.Errorf("TestDBFeedToFeed Unexpected:\n%v\n", output)
				return
			}
		})
	}
}
