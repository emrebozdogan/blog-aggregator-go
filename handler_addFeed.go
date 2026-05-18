package main

import (
	"context"
	"fmt"
	"time"

	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("addfeed command needs to take two arguments")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})

	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("feed is added by %v\n", feedFollow.UserName)
	fmt.Printf("%v is also follows this %v feed", feedFollow.UserName, feedFollow.FeedName)
	fmt.Printf("id: %v\ncreated_at: %v\nupdated_at: %v\nname: %v\nurl: %v\nuser_id: %v\n", feed.ID, feed.CreatedAt, feed.UpdatedAt, feed.Name, feed.Url, feed.UserID)
	return nil

}
