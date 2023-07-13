package main

import "net/http"

func handlerErr(w http.ResponseWriter, req *http.Request) {
    respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}

