package main

import (
	"context"
	"fmt"

	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
)

func handlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("following commmand does not take any arguments")
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Followed feeds by %v\n", user.Name)
	for _, feed := range feedFollows {
		fmt.Printf(" * %v\n", feed.FeedName)
	}

	return nil
}
