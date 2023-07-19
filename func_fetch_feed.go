package main

import (
    "context"
    "fmt"

    _ "github.com/lib/pq"
)

func (cfg *apiConfig) fetchFeed(feed Feed) (RSS, error) {
    ctx := context.Background()

    rss, err := fetchRSS(feed.Url)
    if err != nil {
        err = fmt.Errorf("Error: cfg.fetchFeed: fetchFeedData %v", err)
        return RSS{}, err
    }

    err = cfg.DB.MarkFeedFetched(ctx, feed.ID)
    if err != nil {
        err = fmt.Errorf("Error: cfg.fetchFeed: cfg.DB.MarkFeedFetched: %v", err)
        return rss, err
    }
    return rss, nil
}

