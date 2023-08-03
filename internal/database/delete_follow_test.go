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
	feedID := uuid.MustParse("4ee81dd1-dd4f-4536-b362-b5cd596e9cc8")

	tests := []struct {
		input     CreateFollowParams
		expect    int
		expectErr string
	}{
		{
			// zero values
			expectErr: `pq: insert or update on table "feed_follows" violates foreign key constraint "feed_follows_feed_id_fkey"`,
		},
		{
			// zero values, except keys
			input: CreateFollowParams{
				ID:     id1,
				UserID: userID,
				FeedID: feedID,
			},
			expect:    1,
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
			expect:    1,
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestDeleteFollow Case #%v:", i), func(t *testing.T) {
			_, err := q.CreateFollow(ctx, test.input)
			if err != nil {
				if !strings.Contains(err.Error(), test.expectErr) {
					t.Errorf("Error: %v\n", err)
					return
				}
			}

			resultBeforeDelete, err := q.db.ExecContext(ctx, GET_FOLLOW, test.input.ID)
			if err != nil {
				t.Errorf("Error: Did not delete feedFollow during test: %v", err)
				return
			}
			rowsNumBefore, err := resultBeforeDelete.RowsAffected()
			if err != nil {
				t.Errorf("Error: Unable to get number of rows affected: %v", err)
				return
			}
			if rowsNumBefore != int64(test.expect) {
				t.Errorf("Unexpected: follow not deleted: %v", rowsNumBefore)
				return
			}

			err = q.DeleteFollow(ctx, test.input.ID)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: %v\n", err)
				return
			}

			resultAfterDelete, err := q.db.ExecContext(ctx, GET_FOLLOW, test.input.ID)
			if err != nil {
				t.Errorf("Error: unable to get feedFollow from DB: %v", err)
				return
			}
			rowsNumAfter, err := resultAfterDelete.RowsAffected()
			if err != nil {
				t.Errorf("Error: Unable to get number of rows affected: %v", err)
				return
			}
			if rowsNumAfter != 0 {
				t.Errorf("Unexpected: follow not deleted: %v", rowsNumAfter)
				return
			}
		})
	}
}
