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
	}{{
		input:     "google.com",
		expectURL: "",
		expectErr: "unsupported protocol scheme",
	},
		{
			input:     "https://google.com",
			expectURL: "",
			expectErr: "invalid",
		},
		{
			input:     "https://wagslane.dev/index.xml",
			expectURL: "https://wagslane.dev/index.xml",
			//expectedLink has to be updated every so often to the latest blog post
			expectLink: "https://blog.boot.dev/news/bootdev-beat-2023-08/",
			expectErr:  "not expecting an error",
		},
		{
			input:      "https://blog.boot.dev/index.xml",
			expectURL:  "https://blog.boot.dev/index.xml",
			expectLink: "https://blog.boot.dev/backend/django-for-backend/",
			expectErr:  "not expecting an error",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("Test Case #%v:", i), func(t *testing.T) {
			output, err := fetchRSSDataFromURL(test.input)
			if err != nil && !strings.Contains(err.Error(), test.expectErr) {
				t.Errorf("Unexpected: %v\n", err)
				return
			}

			if output.URL != test.expectURL {
				t.Errorf("Unexpected: %v\n", output.URL)
				return
			}

			if len(output.Channel.Items) > 0 && output.Channel.Items[0].Link != test.expectLink {
				t.Errorf("Unexpected: %v\n", output.Channel.Items[0].Link)
				return
			}
		})
	}
}
