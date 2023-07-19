package main

import (
	"fmt"
    "testing"
)

func TestFetchRSSData(t *testing.T) {
	tests := []struct {
		input  string
        expect RSS
	}{
		{
			input: "https://wagslane.dev/index.xml",
			expect: RSS{},
		},
		//{
		//	input: "https://blog.boot.dev/index.xml",
        //    expect: RSS{},
		//},
    }

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output, err := fetchRSSData(test.input)
            fmt.Println(output.URL)
            fmt.Println(output.Channel.Items[0].Link)
            if err != nil {
				t.Errorf("Unexpected: %v", err)
				return
			}
		})
	}
}

