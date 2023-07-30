package database

import (
	"context"
	"fmt"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func (q *Queries) TestGetNextFeedsToFetch(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		input     int
		expectErr string
	}{
		{
			// zero value
			expectErr: "sql: no rows in result set",
		},
		{
			input:     1,
			expectErr: "not expecting an error",
		},
		{
			input:     5,
			expectErr: "not expecting an error",
		},
		{
			input:     10,
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetFeedByAPIKey Case #%v:", i), func(t *testing.T) {
			output, err := q.GetNextFeedsToFetch(ctx, int32(test.input))
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: %v\n", err)
				return
			}

			if len(output) != test.input {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
