package main

import (
	"time"

	"github.com/google/uuid"

	"github.com/skovranek/rss_aggregator/internal/database"
)

type Post struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt time.Time
	FeedID      uuid.UUID
}

func databasePostToPost(dbPost database.Post) Post {
	return Post{
		ID:            dbPost.ID,
		CreatedAt:     dbPost.CreatedAt,
		UpdatedAt:     dbPost.UpdatedAt,
		Title: dbPost.Title.String,
		Url:           dbPost.Url,
        Description: dbPost.Description.String,
        PublishedAt: dbPost.PublishedAt.Time,
        FeedID: dbPost.FeedID,
    }
}
