package main

import (
	"context"
	"fmt"

	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
)

func handlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("unfollow command takes an url argument")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		return err
	}

	fmt.Printf("%v is unfollowed %v feed\n", user.Name, feed.Name)
	return nil
}
