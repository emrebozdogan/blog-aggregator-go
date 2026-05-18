package main

import (
	"context"
	"fmt"
	"log"
)

func handlerLogin(s *State, cmd Command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("login handler expects a username")
	}

	user, err := s.db.GetUser(context.Background(), cmd.args[0])

	if err != nil {
		log.Fatal("user with that name is not exists")
	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
