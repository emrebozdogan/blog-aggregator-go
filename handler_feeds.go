package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("feeds command does not take any arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("Feed: %v\nUrl: %v\nAdded By: %v\n**************\n", feed.Name, feed.Url, user.Name)
	}

	return nil
}
