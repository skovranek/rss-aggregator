package main

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"

	"github.com/skovranek/rss_aggregator/internal/database"
)

type apiConfig struct {
	DB    *database.Queries
	limit int32
	port  string
}

type UserParams struct {
	Name string `json:"name"`
}

type FeedParams struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type FollowParams struct {
	FeedID string `json:"feed_id"`
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

type Feed struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Url           string
	UserID        uuid.UUID
	LastFetchedAt time.Time
}

type Follow struct {
	ID        uuid.UUID
	FeedID    uuid.UUID
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

type RSS struct {
	URL     string `xml:"url"`
	Channel struct {
		Title         *string `xml:"title"`
		Description   *string `xml:"description"`
		LastBuildDate *string `xml:"lastBuildDate"`
		Items         []Item  `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title       *string `xml:"title"`
	Link        string  `xml:"link"`
	PubDate     *string `xml:"pubDate"`
	Description *string `xml:"description"`
}
