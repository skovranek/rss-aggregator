package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user User) {
	respondWithJSON(w, http.StatusOK, user)
}
