package database

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const GET_FOLLOW = `-- name: GetFollow :one
SELECT id, feed_id, user_id, created_at, updated_at FROM feed_follows
WHERE id = $1
`

func (q *Queries) TestDeleteFollow(t *testing.T) {
	ctx := context.Background()
	id1 := uuid.New()
	id2 := uuid.New()
	now := time.Now().UTC()
	userID := uuid.MustParse("4fb16356-e009-411c-a2b9-58f358b91e0d")
	feedID := uuid.MustParse("0fb2ba16-de86-465a-9a01-5d640fef4d6f")

	tests := []struct {
		createNewFeed bool
		input         CreateFollowParams
		expectErr     string
	}{
		{
			// zero values
			createNewFeed: false,
			expectErr:     `pq: insert or update on table "feeds" violates foreign key constraint "feeds_user_id_fkey`,
		},
		{
			// zero values, except keys
			createNewFeed: true,
			input: CreateFollowParams{
				ID:     id1,
				UserID: userID,
				FeedID: feedID,
			},
			expectErr: "not expecting an error",
		},
		{
			input: CreateFollowParams{
				ID:        id2,
				FeedID:    feedID,
				UserID:    userID,
				CreatedAt: now,
				UpdatedAt: now,
			},
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestDeleteFollow Case #%v:", i), func(t *testing.T) {
			if test.createNewFeed {
				_, err := q.CreateFollow(ctx, test.input)
				if err != nil {
					t.Errorf("Error: %v\n", err)
					return
				}
			}

			err := q.DeleteFollow(ctx, test.input.ID)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: %v\n", err)
				return
			}
		})
	}
}
