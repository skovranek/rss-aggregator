package main

import (
	"github.com/skovranek/rss_aggregator/internal/database"
)

func configure() (*apiConfig, database.CloseDB) {
	port, dbURL, limit := loadEnvVars()

	dbQueries, dbClose := database.InitDB(dbURL)

	return &apiConfig{
		DB:    dbQueries,
		port:  port,
		limit: limit,
	}, dbClose
}
