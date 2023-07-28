package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/skovranek/rss_aggregator/internal/database"
)

const DUPLICATE_URL_KEY_ERR_MSG = `pq: duplicate key value violates unique constraint "posts_url_key"`

func (cfg *apiConfig) createPost(ctx context.Context, feedID uuid.UUID, item Item) error {
	id := uuid.New()
	now := time.Now()

	titleNullStr, err := strPtrToSQLNullStr(item.Title)
	if err != nil {
		return fmt.Errorf("strPtrToSQLNullStr(itme.Title): %v", err)
	}

	descriptionNullStr, err := strPtrToSQLNullStr(item.Description)
	if err != nil {
		return fmt.Errorf("strPtrToSQLNullStr(itme.Description): %v", err)
	}

	publishedAtNullTime, err := strPtrToSQLNullTime(item.PubDate)
	if err != nil {
		return fmt.Errorf("strPtrToSQLNullTime(item.PubDate): %v", err)
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
	if err != nil && !strings.Contains(err.Error(), DUPLICATE_URL_KEY_ERR_MSG) {
		err = fmt.Errorf("unable to add post to database: %v", err)
		return err
	}

	return nil
}

func strPtrToSQLNullStr(strPtr *string) (sql.NullString, error) {
	if strPtr == nil {
		return sql.NullString{}, nil
	}

	nullStr := sql.NullString{}
	err := nullStr.Scan(*strPtr)
	if err != nil {
		err = fmt.Errorf("sql.NullString{}.Scan(*strPtr): %v", err)
		return sql.NullString{}, err
	}
	return nullStr, nil
}

func strPtrToSQLNullTime(strPtr *string) (sql.NullTime, error) {
	if strPtr == nil {
		return sql.NullTime{}, nil
	}

	parsedTime, err := time.Parse(time.RFC1123Z, *strPtr)
	if err != nil {
		err = fmt.Errorf("time.Parse(time.RFC1123Z, *strPtr): %v", err)
		return sql.NullTime{}, err
	}

	nullTime := sql.NullTime{}
	err = nullTime.Scan(parsedTime)
	if err != nil {
		err = fmt.Errorf("sql.NullTime{}.Scan(parsedTime): %v", err)
		return sql.NullTime{}, err
	}
	return nullTime, nil
}
