package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSSData struct {
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
	Link        string `xml:"link"`
	PubDate     *string `xml:"pubDate"`
	Description *string `xml:"description"`
}

func fetchRSSDataFromURL(URL string) (RSSData, error) {
	response, err := http.Get(URL) //#nosec G107
	if err != nil {
		return RSSData{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if response.StatusCode > 299 {
		err = fmt.Errorf("error: response failed. status-code: %d, body: %s", response.StatusCode, body)
		return RSSData{}, err
	}
	if err != nil {
		return RSSData{}, err
	}

	data := RSSData{}
	err = xml.Unmarshal(body, &data)
	if err != nil {
		return RSSData{}, err
	}
	if data.URL == "" {
		data.URL = URL
	}

	return data, nil
}
