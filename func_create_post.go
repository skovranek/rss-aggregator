package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/skovranek/rss_aggregator/internal/database"
)

const DUPLICATE_URL_KEY_ERR_MSG = `pq: duplicate key value violates unique constraint "posts_url_key"`

func (cfg *apiConfig) createPost(ctx context.Context, feedID uuid.UUID, item Item) error {
	now := time.Now()
	id := uuid.New()

    titleNullStr := sql.NullString{}
    err := titleNullStr.Scan(*item.Title)
    if err != nil {
        err = fmt.Errorf("Error: cfg.createPost: titleNullStr.Scan(item.Title): %v", err)
    }

    descriptionNullStr := sql.NullString{}
    err = descriptionNullStr.Scan(*item.Description)
    if err != nil {
        err = fmt.Errorf("Error: cfg.createPost: descriptionNullStr.Scan(item.Description): %v", err)
    }

	publishedAt, err := time.Parse(time.RFC1123Z, *item.PubDate)
	if err != nil {
		err = fmt.Errorf("Error: cfg.createPost: time.Parse(format, item.PubDate): %v", err)
        return err
	}
	publishedAtNullTime := sql.NullTime{}
	err = publishedAtNullTime.Scan(publishedAt)
	if err != nil {
		err = fmt.Errorf("Error: cfg.CreatePost: publishedAtNullTime.Scan(publishedAt): %v", err)
	    return err
    }
	_, err = cfg.DB.CreatePost(ctx, database.CreatePostParams{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		Title:       titleNullStr,
		Url:         item.Link,
		Description: descriptionNullStr,
		PublishedAt: publishedAtNullTime,
		FeedID:      feedID,
	})
	if err != nil && err.Error() != DUPLICATE_URL_KEY_ERR_MSG {
		err = fmt.Errorf("Error: cfg.createPost: cfg.DB.CreatePost: %v", err)
		return err
	}

    return nil
}
