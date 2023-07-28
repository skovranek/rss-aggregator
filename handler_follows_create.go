package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/skovranek/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerFollowsCreate(w http.ResponseWriter, r *http.Request, user User) {
	followParams, err := getFollowParams(r.Body)
	if err != nil {
		log.Printf("Error: handlerFollowsCreate: getFeedParams(r.Body): %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to decode request body")
		return
	}

	feedID := uuid.MustParse(followParams.FeedID)
	userID := user.ID
	id := uuid.New()
	now := time.Now()
	ctx := context.Background()

	follow, err := cfg.DB.CreateFollow(ctx, database.CreateFollowParams{
		ID:        id,
		FeedID:    feedID,
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		log.Printf("Error: handlerFollowsCreate: cfg.DB.CreateFollow: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to add follow to database")
		return
	}

	respondWithJSON(w, http.StatusCreated, follow)
}
