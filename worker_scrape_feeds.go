package main

import (
	"context"
	"log"
	"sync"
	"time"

)

func (cfg *apiConfig) workerScrapeFeeds() {
	wg := sync.WaitGroup{}

	for {
		feeds, err := cfg.getFeedsToFetch()
		if err != nil {
			log.Printf("Error: cfg.workerScrapeFeeds: cfg.getFeedsToFetch: %v", err)
			return
		}

		for _, feed := range feeds {
			wg.Add(1)

			go func(feed Feed) {
				defer wg.Done()

				data, err := fetchRSSDataFromURL(feed.Url)
				if err != nil {
					log.Printf("Error: cfg.workerScrapeFeeds: fetchRSSDataFromURL %v", err)
				}

				ctx := context.Background()
				err = cfg.DB.MarkFeedFetched(ctx, feed.ID)
				if err != nil {
					log.Printf("Error: cfg.workerScrapeFeeds: cfg.DB.MarkFeedFetched: %v", err)
				}

				for _, item := range data.Channel.Items {
				    err = cfg.createPost(ctx, feed.ID, item)
                    if err != nil {
                        log.Printf("Error: cfg.workerScrapeFeeds: cfg.CreatePost(ctx, item, feed.ID): %s", err)
                    }
				}
			}(feed)
		}
		wg.Wait()
		time.Sleep(time.Minute)
	}
}
