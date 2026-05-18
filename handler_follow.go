package main

import (
	"context"
	"fmt"
	"time"

	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("follow command needs to get a url")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollowRecord, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Println("feed follow record is created")
	fmt.Printf("Followed Feed: %v\n Current User Who Followed: %v\n", feedFollowRecord.FeedName, feedFollowRecord.UserName)
	return nil
}
