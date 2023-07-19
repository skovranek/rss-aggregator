package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (cfg *apiConfig) handlerFollowsDelete(w http.ResponseWriter, r *http.Request) {
	followIDStr := chi.URLParam(r, "feedFollowID")
	followID := uuid.MustParse(followIDStr)
	ctx := context.Background()

	err := cfg.DB.DeleteFollow(ctx, followID)
	if err != nil {
		log.Printf("Error: handlerFollowsDelete: cfg.DB.DeleteFollow: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Unable to remove follow from database")
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
