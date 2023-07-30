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

func (q *Queries) TestCreateFollow(t *testing.T) {
	ctx := context.Background()
	id1 := uuid.New()
	id2 := uuid.New()
	now := time.Now().UTC()
	userID := uuid.MustParse("4fb16356-e009-411c-a2b9-58f358b91e0d")
	feedID := uuid.MustParse("0fb2ba16-de86-465a-9a01-5d640fef4d6f")

	tests := []struct {
		input     CreateFollowParams
		expect    FeedFollow
		expectErr string
	}{
		{ // zero values
			expectErr: `pq: insert or update on table "feeds" violates foreign key constraint "feeds_user_id_fkey`,
		},
		{ // zero values, except keys
			input: CreateFollowParams{
				ID:     id1,
				UserID: userID,
				FeedID: feedID,
			},
			expect: FeedFollow{
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
			expect: FeedFollow{
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
		t.Run(fmt.Sprintf("TestCreateFollow Case #%v:", i), func(t *testing.T) {
			output, err := q.CreateFollow(ctx, test.input)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: %v\n", err)
				return
			} else {
				defer func() {
					_, err := q.db.ExecContext(ctx, deleteFollow, output.ID)
					if err != nil {
						t.Errorf("Error: Unable to delete rows from test%v", err)
						return
					}
				}()
			}

			if output.ID != test.expect.ID {
				t.Errorf("Unexpected: id\n%v", output)
				return
			}

			if output.FeedID != test.expect.FeedID {
				t.Errorf("Unexpected: feed_id\n%v", output)
				return
			}

			if output.UserID != test.expect.UserID {
				t.Errorf("Unexpected: user_id\n%v", output)
				return
			}

			if output.CreatedAt.Compare(test.expect.CreatedAt) != 0 {
				t.Errorf("Unexpected: created_at\n%v", output)
				return
			}

			if output.UpdatedAt.Compare(test.expect.UpdatedAt) != 0 {
				t.Errorf("Unexpected: updated_at\n%v", output)
				return
			}

			if output.CreatedAt.Compare(test.expect.UpdatedAt) != 0 {
				t.Errorf("Unexpected: created_at != updated_at\n%v", output)
				return
			}
		})
	}
}
