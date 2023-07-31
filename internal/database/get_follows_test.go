package database

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (q *Queries) TestGetFollowsByUser(t *testing.T) {
	ctx := context.Background()
	userID := uuid.MustParse("4fb16356-e009-411c-a2b9-58f358b91e0d")

	tests := []struct {
		input     uuid.UUID
		expectErr string
	}{
		{
			// zero value
			expectErr: "sql: no rows in result set",
		},
		{
			input:     userID,
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetFollowsByUser Case #%v:", i), func(t *testing.T) {
			rowNotFetched := q.db.QueryRowContext(ctx, GET_FEED_BY_ID, test.input)
			var feedNotFetched Feed
			err := rowNotFetched.Scan(
				&feedNotFetched.ID,
				&feedNotFetched.CreatedAt,
				&feedNotFetched.UpdatedAt,
				&feedNotFetched.Name,
				&feedNotFetched.Url,
				&feedNotFetched.UserID,
				&feedNotFetched.LastFetchedAt,
			)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: Unable to get feed by id: %v", err)
				return
			}

			err = q.GetFollowsByUser(ctx, test.input)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: %v\n", err)
				return
			}

			rowFetched := q.db.QueryRowContext(ctx, GET_FEED_BY_ID, test.input)
			var feedFetched Feed
			err = rowFetched.Scan(
				&feedFetched.ID,
				&feedFetched.CreatedAt,
				&feedFetched.UpdatedAt,
				&feedFetched.Name,
				&feedFetched.Url,
				&feedFetched.UserID,
				&feedFetched.LastFetchedAt,
			)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: Unable to get feed by id: %v", err)
				return
			}

			if feedFetched.ID != test.input || feedFetched.ID != feedNotFetched.ID {
				t.Errorf("Unexpected: id\n%v", feedFetched)
				return
			}

			if !feedFetched.UpdatedAt.After(feedNotFetched.UpdatedAt) {
				t.Errorf("Unexpected: updatedAt\n%v", feedFetched)
				return
			}

			if !feedFetched.LastFetchedAt.Time.After(feedNotFetched.LastFetchedAt.Time) {
				t.Errorf("Unexpected: lastFetchedAt\n%v", feedFetched)
			}
		})
	}
}
