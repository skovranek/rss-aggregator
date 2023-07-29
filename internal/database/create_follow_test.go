package database

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const DELETE_FEED_BY_ID string = `-- name: DeleteFeedByID :exec
DELETE FROM feeds
WHERE id = $1
`

func (q *Queries) TestCreateFeed(t *testing.T) {
	ctx := context.Background()
	id1 := uuid.New()
	id2 := uuid.New()
	now := time.Now().UTC()
	str := "this is a string"
	userID := uuid.MustParse("4fb16356-e009-411c-a2b9-58f358b91e0d")
	nullTime := sql.NullTime{}

	tests := []struct {
		input     CreateFeedParams
		expect    Feed
		expectErr string
	}{
		{ // zero values
			expectErr: `pq: insert or update on table "feeds" violates foreign key constraint "feeds_user_id_fkey`,
		},
		{ // zero values, except keys
			input: CreateFeedParams{
				ID:     id1,
				UserID: userID,
			},
			expect: Feed{
				ID:     id1,
				UserID: userID,
			},
			expectErr: "not expecting an error",
		},
		{
			input: CreateFeedParams{
				ID:        id2,
				CreatedAt: now,
				UpdatedAt: now,
				Name:      str,
				Url:       str,
				UserID:    userID,
			},
			expect: Feed{
				ID:            id2,
				CreatedAt:     now,
				UpdatedAt:     now,
				Name:          str,
				Url:           str,
				UserID:        userID, // check present
				LastFetchedAt: sql.NullTime{},
			},
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCreateFeed Case #%v:", i), func(t *testing.T) {
			output, err := q.CreateFeed(ctx, test.input)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: %v\n", err)
				return
			} else {
				defer func() {
					_, err := q.db.ExecContext(ctx, DELETE_FEED_BY_ID, output.ID)
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

			if output.CreatedAt.Compare(test.expect.CreatedAt) != 0 {
				t.Errorf("Unexpected: createdAt\n%v", output)
				return
			}

			if output.UpdatedAt.Compare(test.expect.UpdatedAt) != 0 {
				t.Errorf("Unexpected: updatedAt\n%v", output)
				return
			}

			if output.Name != test.expect.Name {
				t.Errorf("Unexpected: name\n%v", output)
				return
			}

			if output.Url != test.expect.Url {
				t.Errorf("Unexpected: url\n%v", output)
				return
			}

			if output.UserID != test.expect.UserID {
				t.Errorf("Unexpected: userID\n%v", output)
				return
			}

			// check if an LastFetchedAt was added by DB
			if output.LastFetchedAt != nullTime {
				t.Errorf("Unexpected: lastFetchedAt\n%v", output)
			}
		})
	}
}
