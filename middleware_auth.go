package main

import (
    "context"
    "log"
    "net/http"
    "strings"

    "github.com/skovranek/rss_aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(next authedHandler) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKeyAndPrefix := r.Header.Get("Authorization")
        apiKey := strings.TrimPrefix(apiKeyAndPrefix, "ApiKey ")
        
        ctx := context.Background()

        user, err := cfg.DB.GetUserByAPIKey(ctx, apiKey)
        if err != nil {
            log.Printf("Error: middlewareAuth: cfg.DB.GetUserByAPIKey: %v", err)
            respondWithError(w, http.StatusInternalServerError, "Unable to retrieve user from database")
            return
        }
        next(w, r, user)
    })
}

