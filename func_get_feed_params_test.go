package main

import (
    "encoding/json"
	"fmt"
    "io"
    "strings"
    "testing"
)

type TestFeedParams struct{
    Name string `json:"name"`
    URL string `json:"url"`
}

func TestGetFeedParams(t *testing.T) {
	exampleStr := "this is a string"

	tests := []struct {
        input TestFeedParams
		expect FeedParams
	}{
        {}, // zero values
		{
			input: TestFeedParams{
                Name: exampleStr,
            },
			expect: FeedParams{
                Name: exampleStr,
            },
        },
		{
			input: TestFeedParams{
                URL: exampleStr,
            },
			expect: FeedParams{
                URL: exampleStr,
            },
        },
		{
			input: TestFeedParams{
                Name: exampleStr,
                URL: exampleStr,
            },
			expect: FeedParams{
                Name: exampleStr,
                URL: exampleStr,
            },
        },
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
            b, err := json.Marshal(test.input)
            if err != nil {
                t.Errorf("Error: Test Case #%v: %v", i, err)
                return
            }

            reader := strings.NewReader(string(b))
            readCloser := io.NopCloser(reader)

			output, err := getFeedParams(readCloser)
            if err != nil {
                t.Errorf("Error: Test Case #%v: %v", i, err)
                return
            }

            if output != test.expect {
				t.Errorf("Unexpected:\n%v", output)
				return
			}
		})
	}
}
