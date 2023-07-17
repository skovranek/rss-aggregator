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

    timeStamp := time.Now()
    testStr := "test case"

	tests := []struct {
		input  database.Feed
		expect Feed
	}{
		{
			input: database.Feed{
                ID: id,
                CreatedAt: timeStamp,
                UpdatedAt: timeStamp,
                Name: testStr,
                Url: testStr,
                UserID: id,
                LastFetchedAt: sql.NullTime{
                    Time: timeStamp,
                },
            },
			expect: Feed{
                ID: id,
                CreatedAt: timeStamp,
                UpdatedAt: timeStamp,
                Name: testStr,
                Url: testStr,
                UserID: id,
                LastFetchedAt: timeStamp,
            },
		},
		{
			input: database.Feed{
                ID: id,
                CreatedAt: timeStamp,
                UpdatedAt: timeStamp,
                Name: testStr,
                Url: testStr,
                UserID: id,
            },
            expect: Feed{
                ID: id,
                CreatedAt: timeStamp,
                UpdatedAt: timeStamp,
                Name: testStr,
                Url: testStr,
                UserID: id,
                LastFetchedAt: time.Time{},
            },
		},
    }

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output := databaseFeedToFeed(test.input)
            if output != test.expect {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}

