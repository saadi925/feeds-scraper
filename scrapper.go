package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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
			log.Fatal("error fetching feeds \n Error :", err)
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
	if err != nil {
		log.Print("url to feed err  ", err)
	}
	dateFormats := GetDateFormats()
	for _, item := range rssFeed.Channel.Item {
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
		}
		var publishedAt time.Time
		for _, format := range dateFormats {
			publishedAt, err = time.Parse(format, item.PubDate)
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Printf("could not parse date %v  \n error occured %v ", item.PubDate, err)
			publishedAt = time.Now()
		}
		post, err := db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			ID:          uuid.New(),
			Description: description,
			Url:         item.Link,
			FeedID:      feed.ID,
			PublishedAt: publishedAt,
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("could not create post in db %v  \n error occured %v ", post.Title, err)
			continue
		}
	}
	if err != nil {
		log.Println("err fetching feed", err)
	}
	log.Printf("feed %v collected %d posts", feed.Name, len((rssFeed.Channel.Item)))
}

func GetDateFormats() []string {
	dateFormats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822,
		time.RFC822Z,
	}
	return dateFormats
}
