package main

import (
	"context"
	"fmt"
	"log"
	"os"
)

func handlerReset(s *State, cmd Command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("reset command can't take any arguments")
	}

	if err := s.db.Reset(context.Background()); err != nil {
		log.Fatal("resetting was not successful")
	}

	fmt.Println("resetting was successful")
	os.Exit(0)
	return nil
}
