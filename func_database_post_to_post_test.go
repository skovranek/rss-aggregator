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

	timeNow := time.Now()
	testSQLNullTime := sql.NullTime{
		Time:  timeNow,
		Valid: true,
	}

	testStr := "test case"
	testSQLNullStr := sql.NullString{
		String: testStr,
		Valid:  true,
	}

	tests := []struct {
		input  database.Post
		expect Post
	}{
        {
            input: database.Post{},
            expect: Post{},
        },
		{
			input: database.Post{
				ID:          id,
				CreatedAt:   timeNow,
				UpdatedAt:   timeNow,
				Title:       testSQLNullStr,
				Url:         testStr,
				Description: testSQLNullStr,
				PublishedAt: testSQLNullTime,
				FeedID:      id,
			},
			expect: Post{
				ID:          id,
				CreatedAt:   timeNow,
				UpdatedAt:   timeNow,
				Title:       testStr,
				Url:         testStr,
				Description: testStr,
				PublishedAt: timeNow,
				FeedID:      id,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output := databasePostToPost(test.input)
			if output != test.expect {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
