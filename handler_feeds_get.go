package main

import (
	"context"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	dbFeeds, err := cfg.DB.GetAllFeeds(ctx)
	if err != nil {
		log.Printf("Error: handlerFeedsGet: cfg.DB.GetAllFeeds: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to retrieve feeds from database")
		return
	}

	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feed := databaseFeedToFeed(dbFeed)
		feeds = append(feeds, feed)
	}

	respondWithJSON(w, http.StatusOK, feeds)
}
