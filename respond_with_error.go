package main

import "net/http"

func respondWithError(w http.ResponseWriter, statusCode int, errMsg string) {
    respBody := struct {
        Error string `json:"error"`
    }{
        Error: errMsg,
    }
    respondWithJSON(w, statusCode, respBody)
}

