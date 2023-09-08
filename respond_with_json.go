package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: respondWithJSON: marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if statusCode < 100 || statusCode > 599 {
		log.Printf("Error: respondWithJSON: statusCode out of range: %v", statusCode)
		statusCode = http.StatusConflict
	}
	w.WriteHeader(statusCode)

	_, err = w.Write(data)
	if err != nil {
		log.Printf("Error: respondWithJSON: writing data to response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
