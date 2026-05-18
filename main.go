package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/emrebozdogan/blog-aggregator-go/internal/config"
	"github.com/emrebozdogan/blog-aggregator-go/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	state := State{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := Commands{
		cmds: make(map[string]func(*State, Command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))

	cliArgs := os.Args

	if len(cliArgs) < 2 {
		log.Fatal("requires a command name")
	}

	cmd := Command{
		name: cliArgs[1],
		args: cliArgs[2:],
	}

	if err := commands.run(&state, cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
