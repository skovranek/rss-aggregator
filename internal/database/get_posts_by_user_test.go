package database

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (q *Queries) TestGetPostsByUser(t *testing.T) {
	ctx := context.Background()
	userID := uuid.MustParse("4fb16356-e009-411c-a2b9-58f358b91e0d")

	tests := []struct {
		input      uuid.UUID
		inputLimit int
		expectErr  string
	}{
		{
			// zero value
			expectErr: "sql: no rows in result set",
		},
		{
			input:      userID,
			inputLimit: 10,
			expectErr:  "not expecting an error",
		},
		{
			input:      userID,
			inputLimit: 20,
			expectErr:  "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetPostsByUser Case #%v:", i), func(t *testing.T) {
			getPostsByUserParams := GetPostsByUserParams{
				UserID: test.input,
				Limit:  int32(test.inputLimit),
			}
			outputPosts, err := q.GetPostsByUser(ctx, getPostsByUserParams)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: TestGetPostsByUser: q.GetPostsBuUser: %v\n", err)
				return
			}

			if len(outputPosts) != test.inputLimit {
				t.Errorf("Unexpected: \n%v", len(outputPosts))
				return
			}
		})
	}
}
