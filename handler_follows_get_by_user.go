package main

import (
	"context"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func (cfg *apiConfig) handlerFollowsGetByUser(w http.ResponseWriter, r *http.Request, user User) {
	ctx := context.Background()

	follows, err := cfg.DB.GetFollowsByUser(ctx, user.ID)
	if err != nil {
		log.Printf("Error: handlerFollowsGet: cfg.DB.GetFollows: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to retrieve follows from database")
		return
	}

	respondWithJSON(w, http.StatusOK, follows)
}
