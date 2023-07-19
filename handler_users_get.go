package main

import (
	"net/http"

	"github.com/skovranek/rss_aggregator/internal/database"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, http.StatusOK, user)
}
