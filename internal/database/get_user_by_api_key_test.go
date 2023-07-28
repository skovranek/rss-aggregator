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

func (q *Queries) TestGetUserByAPIKey(t *testing.T) {
    ctx := context.Background()
    id := uuid.MustParse("4fb16356-e009-411c-a2b9-58f358b91e0d")
    timestamp, err := time.Parse(time.RFC3339Nano, "2023-07-11T20:15:24.897281Z")
    if err != nil {
        t.Errorf("Error: %v", err)
    }
    name := "james"
	apiKey := "55a49e10add391404027e65a20156727abedbbcfaafe9316d828c36b4a24a158"

	tests := []struct {
		input  string
		expect User
        expectErr string
	}{
        { 
            // zero value
            expectErr: "sql: no rows in result set",
        },
        {
            input: "not an api key",
            expectErr: "sql: no rows in result set",
        },
		{
			input: apiKey,
			expect: User{
				ID:        id,
				CreatedAt: timestamp,
				UpdatedAt: timestamp,
				Name:      name,
                ApiKey:    apiKey,
            },
            expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetUserByAPIKey Case #%v:", i), func(t *testing.T) {
			output, err := q.GetUserByAPIKey(ctx, test.input)
            if err != nil {
                if strings.Contains(err.Error(), test.expectErr) {
                    return
                }
				t.Errorf("Error: %v\n", err)
				return
			}

            if output.ID != test.expect.ID {
                t.Errorf("Unexpected:\n%v", output)
				return
			}

            if output.CreatedAt.Compare(test.expect.CreatedAt) != 0 {
                t.Errorf("Unexpected:\n%v", output)
				return
			}

            if output.UpdatedAt.Compare(test.expect.UpdatedAt) != 0 {
                t.Errorf("Unexpected:\n%v", output)
				return
            }

            if output.Name != test.expect.Name {
                t.Errorf("Unexpected:\n%v", output)
				return
			}

            // check if an ApiKey was added by DB
            if output.ApiKey == "" {
                t.Errorf("Unexpected:\n%v", output)
            }
		})
	}
}
