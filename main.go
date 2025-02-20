package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/wolv89/bootgator/internal/commands"
	"github.com/wolv89/bootgator/internal/config"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/middleware.go"
	"github.com/wolv89/bootgator/internal/state"

	_ "github.com/lib/pq"
)

func main() {

	appConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	appDB, err := sql.Open("postgres", appConfig.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	appQueries := database.New(appDB)

	appState := state.State{
		Config: &appConfig,
		DB:     appQueries,
	}

	appCommands := commands.Commands{
		List: make(map[string]func(*state.State, commands.Command) error),
	}

	appCommands.Register("login", commands.HandlerLogin)
	appCommands.Register("register", commands.HandlerRegister)
	appCommands.Register("reset", commands.HandlerReset)
	appCommands.Register("users", commands.HandlerUsers)
	appCommands.Register("agg", commands.HandlerAgg)
	appCommands.Register("feeds", commands.HandlerFeeds)

	appCommands.Register("addfeed", middleware.LoggedIn(commands.HandlerAddFeed))
	appCommands.Register("follow", middleware.LoggedIn(commands.HandlerFollow))
	appCommands.Register("unfollow", middleware.LoggedIn(commands.HandlerUnfollow))
	appCommands.Register("following", middleware.LoggedIn(commands.HandlerFollowing))

	if len(os.Args) < 2 {
		log.Fatal("no command given")
	}

	cmd := commands.Command{
		Args: os.Args[2:],
		Name: os.Args[1],
	}

	err = appCommands.Run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
