package main

import (
	"context"

	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		currentUsername := s.cfg.CurrentUserName
		user, err := s.db.GetUser(context.Background(), currentUsername)
		if err != nil {
			return err
		}
		handler(s, cmd, user)
		return nil
	}
}
