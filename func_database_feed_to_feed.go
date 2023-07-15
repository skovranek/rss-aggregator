package main

import (
    "time"

    "github.com/google/uuid"

    "github.com/skovranek/rss_aggregator/internal/database"
)

type Feed struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt time.Time
}

func databaseFeedToFeed(feed database.Feed) Feed {
    return Feed{
        ID: feed.ID,
        CreatedAt: feed.CreatedAt,
        UpdatedAt: feed.UpdatedAt,
        Name: feed.Name,
        Url: feed.Url,
        UserID: feed.UserID,
        LastFetchedAt: feed.LastFetchedAt.Time,
    }
}

