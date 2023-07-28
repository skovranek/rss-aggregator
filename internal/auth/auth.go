package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	apiKeyAndPrefix := header.Get("Authorization")
	if apiKeyAndPrefix == "" {
		return "", errors.New("no authorization header")
	}

	if !strings.Contains(apiKeyAndPrefix, "ApiKey ") {
		return "", errors.New("incorrect formatting of authorization header")
	}

	apiKey := strings.TrimPrefix(apiKeyAndPrefix, "ApiKey ")
	if apiKey == "" {
		return "", errors.New("no api key in authorization header")
	}
	return apiKey, nil
}
