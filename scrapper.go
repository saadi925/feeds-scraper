package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/saadi925/rssagregator/internal/database"
)

func startScrapping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("scrapping on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Fatal("error fetching feeds", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			fmt.Println(feed.ID)
			go scrapeFeed(wg, db, feed)
		}
		wg.Wait()

	}
}
func scrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkedFeedAsFetch(context.Background(), feed.ID)
	if err != nil {
		log.Println("err marking feed as fetched", err)
	}
	rssFeed, err := urlToFeed(feed.Url)
	for _, item := range rssFeed.Channel.Item {
		log.Println("post found ", item.Title)
	}
	if err != nil {
		log.Println("err fetching feed", err)
	}
	log.Printf("collected %v posts found from feed %v ", feed.Name, len((rssFeed.Channel.Item)))
}
