package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
)

type TestFeedParams struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func TestGetFeedParams(t *testing.T) {
	str := "this is a string"

	tests := []struct {
		input  TestFeedParams
		expect FeedParams
	}{
		{}, // zero values
		{
			input: TestFeedParams{
				Name: str,
			},
			expect: FeedParams{
				Name: str,
			},
		},
		{
			input: TestFeedParams{
				URL: str,
			},
			expect: FeedParams{
				URL: str,
			},
		},
		{
			input: TestFeedParams{
				Name: str,
				URL:  str,
			},
			expect: FeedParams{
				Name: str,
				URL:  str,
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestGetFeedParams Case #%v:", i), func(t *testing.T) {
			b, err := json.Marshal(test.input)
			if err != nil {
				t.Errorf("Error: TestGetFeedParams Case #%v: %v", i, err)
				return
			}

			reader := strings.NewReader(string(b))
			readCloser := io.NopCloser(reader)

			output, err := getFeedParams(readCloser)
			if err != nil {
				t.Errorf("Error: TestGetFeedParams Case #%v: %v", i, err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetFeedParams: \n%v", output)
				return
			}
		})
	}
}
