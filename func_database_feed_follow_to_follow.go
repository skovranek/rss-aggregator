package main

import (
	"github.com/skovranek/rss_aggregator/internal/database"
)

func databaseFollowToFollow(dbFollow database.FeedFollow) Follow {
	return Follow{
		ID:        dbFollow.ID,
		FeedID:    dbFollow.FeedID,
		UserID:    dbFollow.UserID,
		CreatedAt: dbFollow.CreatedAt,
		UpdatedAt: dbFollow.UpdatedAt,
	}
}
