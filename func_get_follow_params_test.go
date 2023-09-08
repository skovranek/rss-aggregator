package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
)

type TestFollowParams struct {
	FeedID string `json:"feed_id"`
}

func TestGetFollowParams(t *testing.T) {
	str := "this is a string"

	tests := []struct {
		input  TestFollowParams
		expect FollowParams
	}{
		{}, // zero values
		{
			input: TestFollowParams{
				FeedID: str,
			},
			expect: FollowParams{
				FeedID: str,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetFollowParams Case #%v:", i), func(t *testing.T) {
			b, err := json.Marshal(test.input)
			if err != nil {
				t.Errorf("Error: TestGetFollowParams: Case #%v: %v", i, err)
				return
			}

			reader := strings.NewReader(string(b))
			readCloser := io.NopCloser(reader)

			output, err := getFollowParams(readCloser)
			if err != nil {
				t.Errorf("Error: TestGetFollowParams: Case #%v: %v", i, err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetFollowParams:\n%v", output)
				return
			}
		})
	}
}
