package main

import (
	"context"
	"log"
	"net/http"

	"github.com/skovranek/rss_aggregator/internal/auth"
)

type authedHandler func(http.ResponseWriter, *http.Request, User)

func (cfg *apiConfig) middlewareAuth(next authedHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			log.Printf("Error: middlewareAuth: auth.GetAPIKey: %v", err)
			respondWithError(w, http.StatusUnauthorized, "Unable to parse API Key from Authorization header")
			return
		}

		ctx := context.Background()

		dbUser, err := cfg.DB.GetUserByAPIKey(ctx, apiKey)
		if err != nil {
			log.Printf("Error: middlewareAuth: cfg.DB.GetUserByAPIKey: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Unable to retrieve user from database")
			return
		}

		user := databaseUserToUser(dbUser)

		next(w, r, user)
	})
}
