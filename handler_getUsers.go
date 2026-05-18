package main

import (
	"context"
	"fmt"
	"log"
)

func handlerGetUsers(s *State, cmd Command) error {

	if len(cmd.args) != 0 {
		return fmt.Errorf("users command can't take any arguments")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		log.Fatal("can't get the users right now")
	}

	for _, user := range users {
		if user.Name != s.cfg.CurrentUserName {
			fmt.Printf("* %v\n", user.Name)
			continue
		}
		fmt.Printf("* %v (current)\n", user.Name)
	}

	return nil
}
