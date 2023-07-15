package main

import (
    "context"
	"encoding/json"
    "log"
    "net/http"
    "time"

    "github.com/google/uuid"
    _ "github.com/lib/pq"
    "github.com/skovranek/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
    decoder := json.NewDecoder(r.Body)
	feedParams := struct {
        Name string `json:"name"`
        URL string `json:"url"`
    }{}

	err := decoder.Decode(&feedParams)
	if err != nil {
        log.Printf("Error: handlerFeedsCreate: decoder.Decode(&feedParams): %w", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to decode request body")
		return
	}
    
    name := feedParams.Name
    url := feedParams.URL
    id := uuid.New()
    now := time.Now()
    userID := user.ID
    ctx := context.Background()

    feed, err := cfg.DB.CreateFeed(ctx, database.CreateFeedParams{
        ID: id,
        CreatedAt: now,
        UpdatedAt: now,
        Name: name,
        Url: url,
        UserID: userID,
    }) 
	if err != nil {
        log.Printf("Error: handlerFeedsCreate: database.CreateFeed: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to add feed to database")
		return
	}

    followID := uuid.New()

    follow, err := cfg.DB.CreateFollow(ctx, database.CreateFollowParams{
        ID: followID,
        FeedID: feed.ID,
        UserID: userID,
        CreatedAt: now,
        UpdatedAt: now,
    }) 
	if err != nil {
        log.Printf("Error: handlerFeedsCreate: cfg.DB.CreateFollow: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to add follow to database")
		return
	}

    respBody := struct {
        Feed database.Feed `json:"feed"`
        Follow database.FeedFollow `json:"feed_follow"`
    }{
        Feed: feed,
        Follow: follow,
    }

	respondWithJSON(w, http.StatusCreated, respBody)
}

