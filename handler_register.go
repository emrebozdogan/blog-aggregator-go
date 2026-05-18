package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *State, cmd Command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("register handler expects a username")
	}

	newUser, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})

	if err != nil {
		log.Fatal("user with that name already exists")
	}

	if err = s.cfg.SetUser(newUser.Name); err != nil {
		return err
	}

	fmt.Println("user is created")
	fmt.Printf("id: %v\ncreated_at: %v\nupdated_at: %v\nname: %v\n", newUser.ID, newUser.CreatedAt, newUser.UpdatedAt, newUser.Name)
	return nil
}
