package database

import (
    "context"
	"fmt"
    "strings"
	"testing"

    _ "github.com/lib/pq"
)

func (q *Queries) TestGetAllFeeds(t *testing.T) {
    ctx := context.Background()

	tests := []struct {
		input  context.Context
        expectErr string
	}{
        {
            // zero value for ctx
            expectErr: "sql: no rows in result set",
        },
        {
            input: ctx,
            expectErr: "not expecting an error",
        },
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetAllFeeds Case #%v:", i), func(t *testing.T) {
			output, err := q.GetAllFeeds(ctx)
            if err != nil {
                if strings.Contains(err.Error(), test.expectErr) {
                    return
                }
				t.Errorf("Error: %v\n", err)
				return
			}

            if output != nil {
                t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
