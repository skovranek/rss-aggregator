package main

import (
    "context"
    "log"
    "net/http"

    _ "github.com/lib/pq"
)

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()

    feeds, err := cfg.DB.GetAllFeeds(ctx)
    if err != nil {
        log.Printf("Error: handlerFeedsGet: cfg.DB.GetAllFeeds: %v", err)
        respondWithError(w, http.StatusInternalServerError, "Unable to retrieve feeds from database")
        return
    }

	respondWithJSON(w, http.StatusOK, feeds)
}

