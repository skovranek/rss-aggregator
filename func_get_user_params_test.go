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
	str := "this is a string"

	tests := []struct {
		input  TestUserParams
		expect UserParams
	}{
		{}, // zero values
		{
			input: TestUserParams{
				Name: str,
			},
			expect: UserParams{
				Name: str,
			},
		},
	}

	for i, test := range tests {
        t.Run(fmt.Sprintf("TestGetUserParams Case #%v:", i), func(t *testing.T) {
			b, err := json.Marshal(test.input)
			if err != nil {
                t.Errorf("Error: TestGetUserParams Case #%v: %v", i, err)
				return
			}

			reader := strings.NewReader(string(b))
			readCloser := io.NopCloser(reader)

			output, err := getUserParams(readCloser)
			if err != nil {
				t.Errorf("Error: TestGetUserParams Case #%v: %v", i, err)
				return
			}

			if output != test.expect {
                t.Errorf("Unexpected: TestGetUserParams:\n%v", output)
				return
			}
		})
	}
}
