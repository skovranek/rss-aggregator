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

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userParams := struct {
        Name string `json:"name"`
    }{}

	err := decoder.Decode(&userParams)
	if err != nil {
        log.Printf("Error: handlerUsersCreate: decoder.Decode(%userParams): %w", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to decode request body")
		return
	}
    
    name := userParams.Name
    id := uuid.New()
    now := time.Now()

    ctx := context.Background()

    user, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
        ID: id,
        CreatedAt: now,
        UpdatedAt: now,
        Name: name,
    }) 
	if err != nil {
        log.Printf("Error: handlerUsersCreate: cfg.DB.CreateUser: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to add user to database")
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}
