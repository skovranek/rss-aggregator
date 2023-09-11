package main

import (
	"github.com/skovranek/rss_aggregator/internal/database"
)

func databasePostToPost(dbPost database.Post) Post {
	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title.String,
		Url:         dbPost.Url,
		Description: dbPost.Description.String,
		PublishedAt: dbPost.PublishedAt.Time,
		FeedID:      dbPost.FeedID,
	}
}
