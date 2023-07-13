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
		log.Printf("Error decoding request parameters: %w", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't get request body")
		return
	}
    
    name := userParams.Name
    id := uuid.New()
    createdAt := time.Now()

    ctx := context.Background()

    insertedUser, err := cfg.DB.CreateUser(ctx, database.CreateUserParams{
        ID: id,
        CreatedAt: createdAt,
        UpdatedAt: createdAt,
        Name: name,
    }) 
	if err != nil {
		log.Printf("Error creating user in database: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Coundn't create user in database")
		return
	}

	respondWithJSON(w, http.StatusCreated, insertedUser)
}

