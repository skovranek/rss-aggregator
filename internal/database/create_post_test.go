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

const DELETE_POST_BY_ID string = `-- name: DeletePostByID :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) TestCreatePost(t *testing.T) {
	ctx := context.Background()
	id1 := uuid.New()
	id2 := uuid.New()
	now := time.Now().UTC()
	removedMonotonicClockTimeStr := now.Format(time.Layout)
	now, err := time.Parse(time.Layout, removedMonotonicClockTimeStr)
	if err != nil {
		t.Errorf("Error: TestCreatePost: time.Parse() unable to parse time test variable: %v", err)
	}
	str1 := "this is a string 1"
	str2 := "this is a string 2"
	str3 := "this is a string 3"
	nullStr := sql.NullString{}
	err = nullStr.Scan(str3)
	if err != nil {
		t.Errorf("Error: TestCreatePost: sql.NullString{}.Scan(str): %v", err)
		return
	}
	nullTimeNow := sql.NullTime{}
	err = nullTimeNow.Scan(now)
	if err != nil {
		t.Errorf("Error: TestCreatePost: sql.NullTime{}.Scan(now): %v", err)
		return
	}
	feedID := uuid.MustParse("4ee81dd1-dd4f-4536-b362-b5cd596e9cc8")

	tests := []struct {
		input     CreatePostParams
		expect    Post
		expectErr string
	}{
		{
			// zero values
			expectErr: `pq: insert or update on table "posts" violates foreign key constraint "posts_feed_id_fkey"`,
		},
		{
			// zero values, except keys
			input: CreatePostParams{
				ID:     id1,
				Url:    str1,
				FeedID: feedID,
			},
			expect: Post{
				ID:     id1,
				Url:    str1,
				FeedID: feedID,
			},
			expectErr: "not expecting an error",
		},
		{
			input: CreatePostParams{
				ID:          id2,
				CreatedAt:   now,
				UpdatedAt:   now,
				Title:       nullStr,
				Url:         str2,
				Description: nullStr,
				PublishedAt: nullTimeNow,
				FeedID:      feedID,
			},
			expect: Post{
				ID:          id2,
				CreatedAt:   now,
				UpdatedAt:   now,
				Title:       nullStr,
				Url:         str2,
				Description: nullStr,
				PublishedAt: nullTimeNow,
				FeedID:      feedID,
			},
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCreatePost Case #%v:", i), func(t *testing.T) {
			output, err := q.CreatePost(ctx, test.input)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: TestCreatePost: %v", err)
				return
			} else {
				defer func() {
					_, err := q.db.ExecContext(ctx, DELETE_POST_BY_ID, output.ID)
					if err != nil {
						t.Errorf("Error: TestCreatePost: Unable to delete rows from test: %v", err)
						return
					}
				}()
			}

			if output.ID != test.expect.ID {
				t.Errorf("Unexpected: id\n%v", output)
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

			if output.Title != test.expect.Title {
				t.Errorf("Unexpected: title\n%v", output)
				return
			}

			if output.Url != test.expect.Url {
				t.Errorf("Unexpected: url\n%v", output)
				return
			}

			if output.Description != test.expect.Description {
				t.Errorf("Unexpected: description\n%v", output)
				return
			}

			if output.PublishedAt != test.expect.PublishedAt {
				t.Errorf("Unexpected: published_at\n%v\n%v", output.PublishedAt, test.expect.PublishedAt)
				return
			}

			if output.FeedID != test.expect.FeedID {
				t.Errorf("Unexpected: feed_id\n%v", output)
				return
			}
		})
	}
}
