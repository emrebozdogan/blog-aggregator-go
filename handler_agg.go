package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *State, cmd Command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("agg needs to take time between request argument")
	}
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
