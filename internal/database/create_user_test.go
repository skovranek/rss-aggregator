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

const DELETE_USER_BY_ID string = `-- name: DeleteUserByID :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) TestCreateUser(t *testing.T) {
    ctx := context.Background()
	id1 := uuid.New()
    id2 := uuid.New()
	now := time.Now().UTC()
    str := "this is a string"

	tests := []struct {
		input  CreateUserParams
		expect User
        expectErr string
	}{
        { // zero values, "" is used as primary key user_id
            expectErr: `pq: duplicate key value violates unique constraint "users_pkey"`,
        },
        { // zero values, except primary key user_id
            input: CreateUserParams{
                ID: id1,
            },
            expect: User{
                ID: id1,
            },
            expectErr: "not expecting an error",
        },
		{
			input: CreateUserParams{
				ID:        id2,
                CreatedAt: now,
				UpdatedAt: now,
				Name:      str,
			},
			expect: User{
				ID:        id2,
				CreatedAt: now,
				UpdatedAt: now,
				Name:      str,
                ApiKey:    str, // check present
			},
            expectErr: "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCreateUser Case #%v:", i), func(t *testing.T) {
			output, err := q.CreateUser(ctx, test.input)
            if err != nil {
                if strings.Contains(err.Error(), test.expectErr) {
                    return
                }
                t.Errorf("Error: %v\n", err)
                return
			} else {
                defer func() {
                    _, err := q.db.ExecContext(ctx, DELETE_USER_BY_ID, output.ID) 
                    if err != nil {
                        t.Errorf("Error: Unable to delete rows from test%v", err)
                        return
                    }
                }()
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
