package main

import (
    "context"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    databaseFeeds, err := cfg.DB.GetAllFeeds(ctx)
    if err != nil {
        log.Printf("Error: handlerFeedsGet: cfg.DB.GetAllFeeds: %v", err)
        respondWithError(w, http.StatusInternalServerError, "Unable to retrieve feeds from database")
        return
    }

    feeds := []Feed{}
    for _, databaseFeed := range databaseFeeds {
        feed := databaseFeedToFeed(databaseFeed)
        feeds = append(feeds, feed)
    }

	respondWithJSON(w, http.StatusOK, feeds)
}

