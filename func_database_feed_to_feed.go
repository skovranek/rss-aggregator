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

func databaseFeedToFeed(dbFeed database.Feed) Feed {
    return Feed{
        ID: dbFeed.ID,
        CreatedAt: dbFeed.CreatedAt,
        UpdatedAt: dbFeed.UpdatedAt,
        Name: dbFeed.Name,
        Url: dbFeed.Url,
        UserID: dbFeed.UserID,
        LastFetchedAt: dbFeed.LastFetchedAt.Time,
    }
}

