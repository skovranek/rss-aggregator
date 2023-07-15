package main

import (
    "context"
    "log"
    "net/http"

    _ "github.com/lib/pq"

    "github.com/skovranek/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
    ctx := context.Background()

    follows, err := cfg.DB.GetFollows(ctx, user.ID)
    if err != nil {
        log.Printf("Error: handlerFollowsGet: cfg.DB.GetFollows: %v", err)
        respondWithError(w, http.StatusInternalServerError, "Unable to retrieve follows from database")
        return
    }

	respondWithJSON(w, http.StatusOK, follows)
}

