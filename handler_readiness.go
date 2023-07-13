package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
    respBody := struct {
        Status string `json:"status"`
    }{
        Status: "ok",
    }
    respondWithJSON(w, http.StatusOK, respBody)
}

