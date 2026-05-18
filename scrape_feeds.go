package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/emrebozdogan/blog-aggregator-go/internal/api"
	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *State) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	err = s.db.MarkFeedFetched(context.Background(), feedToFetch.ID)
	if err != nil {
		return err
	}

	rssFeed, err := api.FetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		return err
	}

	for _, post := range rssFeed.Channel.Item {
		pubDate := sql.NullTime{Valid: false}
		t, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", post.PubDate)
		if err == nil {
			pubDate = sql.NullTime{Time: t, Valid: true}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       post.Title,
			Url:         post.Link,
			Description: post.Description,
			PublishedAt: pubDate,
			FeedID:      feedToFetch.ID,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				log.Printf("error creating post: %v", err)
			}
		}
	}
	return nil
}
