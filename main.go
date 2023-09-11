package main

import (
	"log"
	// "time"
)

func main() {
	cfg, dbClose := configure()

	defer dbClose()

	//	go cfg.scrapeFeeds(time.Minute)

	srv := cfg.createServer()

	log.Printf("Starting server on port #%s", cfg.port)
	log.Fatal(srv.ListenAndServe())
}
