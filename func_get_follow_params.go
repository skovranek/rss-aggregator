package main

import (
	"encoding/json"
	"fmt"
	"io"
)

func getFollowParams(readCloser io.ReadCloser) (FollowParams, error) {
	decoder := json.NewDecoder(readCloser)
	followParams := FollowParams{}

	err := decoder.Decode(&followParams)
	if err != nil {
		err = fmt.Errorf("unable to decode request body and get feedFollow parameters: %v", err)
		return FollowParams{}, err
	}

	return followParams, nil
}
