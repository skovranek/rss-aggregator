package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type FeedParams struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getFeedParams(readCloser io.ReadCloser) (FeedParams, error) {
	decoder := json.NewDecoder(readCloser)
	feedParams := FeedParams{}

	err := decoder.Decode(&feedParams)
	if err != nil {
        err = fmt.Errorf("unable to decode request body and get feed parameters: %v", err)
		return FeedParams{}, err
	}

	return feedParams, nil
}
