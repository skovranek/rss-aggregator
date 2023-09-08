package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestFetchRSSDataFromURL(t *testing.T) {
	tests := []struct {
		input      string
		expectURL  string
		expectLink string
		expectErr  string
	}{
		//{
		//	input:     "google.com",
		//	expectErr: "unsupported protocol scheme",
		//},
		//{
        //    input:     "https://google.com",
		//	expectErr: "invalid",
		//},
		{
			input:     "https://wagslane.dev/index.xml",
			expectURL: "https://wagslane.dev/index.xml",
			//expectedLink has to be updated every so often to the latest blog post
			expectLink: "https://wagslane.dev/about/",
			expectErr:  "not expecting an error",
		},
		{
			input:      "https://blog.boot.dev/index.xml",
			expectURL:  "https://blog.boot.dev/index.xml",
            expectLink: "https://blog.boot.dev/cryptography/node-js-random-number/",
			expectErr:  "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestFetchRSSDataFromURLTest Case #%v:", i), func(t *testing.T) {
			output, err := fetchRSSFromURL(test.input)
			if err != nil && !strings.Contains(err.Error(), test.expectErr) {
                t.Errorf("Unexpected: TestFetchRSSDataFromURL: %v", err)
				return
			}

			if output.URL != test.expectURL {
                t.Errorf("Unexpected: TestFetchRSSDataFromURL: %v", output.URL)
				return
			}

            itemsLen := len(output.Channel.Items)
			if itemsLen > 0 && output.Channel.Items[itemsLen-1].Link != test.expectLink {
                t.Errorf("Unexpected: TestFetchRSSDataFromURL: %v", output.Channel.Items[itemsLen-1].Link)
				return
			}
		})
	}
}
