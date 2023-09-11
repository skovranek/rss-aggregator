package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

func fetchRSSFromURL(URL string) (RSS, error) {
	response, err := http.Get(URL) //#nosec G107
	if err != nil {
		return RSS{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if response.StatusCode > 299 {
		err = fmt.Errorf("response failure: status code: %d, body: %s", response.StatusCode, body)
		return RSS{}, err
	}
	if err != nil {
		return RSS{}, err
	}

	rss := RSS{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return RSS{}, err
	}
	if rss.URL == "" {
		rss.URL = URL
	}

	return rss, nil
}
