package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func (cfg *apiConfig) workerFetchFeeds() {
	wg := sync.WaitGroup{}

	for {
		feeds, err := cfg.getFeedsToFetch()
		if err != nil {
			log.Printf("Error: cfg.workerFetchFeeds: cfg.getFeedsToFetch: %v", err)
			return
		}

		for _, feed := range feeds {
			wg.Add(1)

			go func(feed Feed) {
				defer wg.Done()

				data, err := fetchRSSDataFromURL(feed.Url)
				if err != nil {
					log.Printf("Error: cfg.workerFetchFeeds: fetchRSSDataFromURL %v", err)
				}

				ctx := context.Background()
				err = cfg.DB.MarkFeedFetched(ctx, feed.ID)
				if err != nil {
					log.Printf("Error: cfg.workerFetchFeeds: cfg.DB.MarkFeedFetched: %v", err)
				}

				for _, item := range data.Channel.Items {
					fmt.Println(item.Title)
				}
			}(feed)
		}
		wg.Wait()
		time.Sleep(time.Minute)
	}
}
