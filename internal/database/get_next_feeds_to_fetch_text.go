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
		input     int32
		expect    int
		expectErr string
	}{
		{
			// zero value
			expectErr: "sql: no rows in result set",
		},
		{
			input:     int32(1),
			expect:    1,
			expectErr: "not expecting an error",
		},
		{
			input:     int32(5),
			expect:    5,
			expectErr: "not expecting an error",
		},
		{
			input:     int32(10),
			expect:    10,
			expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetFeedByAPIKey Case #%v:", i), func(t *testing.T) {
			output, err := q.GetNextFeedsToFetch(ctx, test.input)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Error: %v\n", err)
				return
			}

			if len(output) != test.expect {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
