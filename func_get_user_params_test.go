package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"testing"
)

type TestUserParams struct {
	Name string `json:"name"`
}

func TestGetUserParams(t *testing.T) {
	exampleStr := "this is a string"

	tests := []struct {
		input  TestUserParams
		expect UserParams
	}{
		{}, // zero values
		{
			input: TestUserParams{
				Name: exampleStr,
			},
			expect: UserParams{
				Name: exampleStr,
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

			output, err := getUserParams(readCloser)
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
