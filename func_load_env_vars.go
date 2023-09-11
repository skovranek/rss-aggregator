package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadEnvVars() (string, string, int32) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(`Error: func_load_env_vars.go: godotenv.Load(): cannot load ".env" file`)
	}

	port := os.Getenv("PORT")
	dbURL := os.Getenv("CONN")

	limitStr := os.Getenv("POSTS_PER_FETCH_LIMIT")
	limitInt, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Fatal(`Error: func_load_env_vars.go: strconv.Atoi(limitStr): cannot convert str to int`)
	}
	limit := int32(limitInt) //#nosec G109

	return port, dbURL, limit
}
