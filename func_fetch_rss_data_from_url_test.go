package main

import (
	"fmt"
	"testing"
)

func TestFetchRSSDataFromURL(t *testing.T) {
	tests := []struct {
		input  string
		expect RSSData
	}{
		{
			input:  "https://wagslane.dev/index.xml",
			expect: RSSData{},
		},
		//{
		//	input: "https://blog.boot.dev/index.xml",
		//    expect: RSSData{},
		//},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output, err := fetchRSSDataFromURL(test.input)
			fmt.Println(output.URL)
			fmt.Println(output.Channel.Items[0].Link)
			if err != nil {
				t.Errorf("Unexpected: %v", err)
				return
			}
		})
	}
}
