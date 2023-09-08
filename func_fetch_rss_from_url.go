package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSS struct {
	URL     string `xml:"url"`
	Channel struct {
		Title         *string `xml:"title"`
		Description   *string `xml:"description"`
		LastBuildDate *string `xml:"lastBuildDate"`
		Items         []Item  `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title       *string `xml:"title"`
	Link        string  `xml:"link"`
	PubDate     *string `xml:"pubDate"`
	Description *string `xml:"description"`
}

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
