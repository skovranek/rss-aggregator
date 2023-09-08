package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/skovranek/rss_aggregator/internal/database"
)

func TestDBUserToUser(t *testing.T) {
	id := uuid.New()
	now := time.Now().UTC()
	str := "this is a string"

	tests := []struct {
		input  database.User
		expect User
	}{
		{
			input:  database.User{},
			expect: User{},
		},
		{
			input: database.User{
				ID:        id,
				CreatedAt: now,
				UpdatedAt: now,
				Name:      str,
				ApiKey:    str,
			},
			expect: User{
				ID:        id,
				CreatedAt: now,
				UpdatedAt: now,
				Name:      str,
				ApiKey:    str,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestDBUserToUser Case #%v:", i), func(t *testing.T) {
			output := databaseUserToUser(test.input)
			if output != test.expect {
                t.Errorf("Unexpected: TestDBUserToUser:\n%v", output)
				return
			}
		})
	}
}
