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
	timeNow := time.Now()
	testStr := "test case"

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
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				Name:      testStr,
				ApiKey:    testStr,
			},
			expect: User{
				ID:        id,
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
				Name:      testStr,
				ApiKey:    testStr,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output := databaseUserToUser(test.input)
			if output != test.expect {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
