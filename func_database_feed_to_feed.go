package main

import (
	"github.com/skovranek/rss_aggregator/internal/database"
)

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:            dbFeed.ID,
		CreatedAt:     dbFeed.CreatedAt,
		UpdatedAt:     dbFeed.UpdatedAt,
		Name:          dbFeed.Name,
		Url:           dbFeed.Url,
		UserID:        dbFeed.UserID,
		LastFetchedAt: dbFeed.LastFetchedAt.Time,
	}
}
