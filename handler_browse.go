package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
)

func handlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		n, err := strconv.Atoi(cmd.args[0])
		if err == nil {
			limit = n
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return err
	}

	fmt.Println("Posts for user:")
	for _, post := range posts {
		fmt.Printf(" * %v\n", post.Title)
	}

	return nil
}
