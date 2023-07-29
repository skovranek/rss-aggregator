package main

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"
)

const ERR_MSG_DUPLICATE_URL_KEY = `pq: duplicate key value violates unique constraint "posts_url_key"`

func (cfg *apiConfig) workerScrapeFeeds(interval time.Duration) {
    log.Printf("Starting RSS feed scraper...")

	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
        log.Printf("Scraping...")

		feeds, err := cfg.getFeedsToFetch()
		if err != nil {
			log.Printf("Error: cfg.workerScrapeFeeds: cfg.getFeedsToFetch: %v", err)
			return
		}

		wg := sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go func(feed Feed) {
				defer wg.Done()

                log.Printf("Scraping: %s", feed.Url)

				data, err := fetchRSSDataFromURL(feed.Url)
				if err != nil {
					log.Printf("Error: cfg.workerScrapeFeeds: fetchRSSDataFromURL %v", err)
				}

                ctx := context.Background()
                err = cfg.DB.MarkFeedFetched(ctx, feed.ID)
                if err != nil {
                    log.Printf("Error: cfg.workerScrapeFeeds: cfg.DB.MarkFeedFetched: %v", err)
                }

                log.Printf("Posting from: %s - %s", feed.Url, *data.Channel.Title)

				for _, item := range data.Channel.Items {
					err = cfg.createPost(ctx, feed.ID, item)
					if err != nil {
                        if strings.Contains(err.Error(), ERR_MSG_DUPLICATE_URL_KEY) {
                            continue
                        }
						log.Printf("Error: cfg.workerScrapeFeeds: cfg.CreatePost(ctx, item, feed.ID): %v", err)
					} else {
                        log.Printf("Creating post: %s - %s", *data.Channel.Title, *item.Title)
                    }
				}
			}(feed)
		}
		wg.Wait()
	}
}
