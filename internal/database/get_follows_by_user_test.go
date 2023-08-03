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
		expect    int
		expectErr string
	}{
		{
			// zero value
			expectErr: "sql: no rows in result set",
		},
		{
			input:     userID,
			expect:    1,
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetFollowsByUser Case #%v:", i), func(t *testing.T) {
			outputFollows, err := q.GetFollowsByUser(ctx, test.input)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: TestGetFollowsByUser: q.GetFollowsByUser: %v\n", err)
				return
			}

			if len(outputFollows) != test.expect {
				t.Errorf("Unexpected: \n%v", len(outputFollows))
				return
			}
		})
	}
}
