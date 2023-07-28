package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type UserParams struct {
	Name string `json:"name"`
}

func getUserParams(readCloser io.ReadCloser) (UserParams, error) {
	decoder := json.NewDecoder(readCloser)
	userParams := UserParams{}

	err := decoder.Decode(&userParams)
	if err != nil {
		err = fmt.Errorf("unable to decode request body and get user parameters: %v", err)
		return UserParams{}, err
	}

	return userParams, nil
}
