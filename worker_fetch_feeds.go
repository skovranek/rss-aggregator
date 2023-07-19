package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

func (cfg *apiConfig) workerFetchFeeds() {
    wg := sync.WaitGroup{}

    for ;; {
        feeds, err := cfg.getFeedsToFetch()
        if err != nil {
            log.Printf("Error: cfg.workerFetchFeeds: cfg.getFeedsToFetch: %v", err)
            return
        }

        for _, feed := range feeds {
            wg.Add(1)

            go func(feed Feed) {
                defer wg.Done()

                rss, err := cfg.fetchFeed(feed)
                if err != nil {
                    log.Printf("Error: cfg.workerFetchFeeds: cfg.fetchFeed: %v", err)
                }

                for _, item := range rss.Channel.Items {
                    fmt.Println(item.Title)
                }
            }(feed)
        }
        wg.Wait()
        time.Sleep(time.Minute)
    }
}

