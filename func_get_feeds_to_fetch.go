package main

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func (cfg *apiConfig) getFeedsToFetch() ([]Feed, error) {
	ctx := context.Background()
	dbFeeds, err := cfg.DB.GetNextFeedsToFetch(ctx, cfg.Limit)
	if err != nil {
		err = fmt.Errorf("Error: cfg.getFeedsToFetch: cfg.DB.GetNextFeedsToFetch: %v", err)
		return []Feed{}, err
	}

	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feed := databaseFeedToFeed(dbFeed)
		feeds = append(feeds, feed)
	}

	return feeds, nil
}
