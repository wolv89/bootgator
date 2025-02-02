package main

import (
	"log"
	"os"

	"github.com/wolv89/bootgator/internal/commands"
	"github.com/wolv89/bootgator/internal/config"
	"github.com/wolv89/bootgator/internal/state"
)

func main() {

	appConfig, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	appState := state.State{
		Config: &appConfig,
	}

	appCommands := commands.Commands{
		List: make(map[string]func(*state.State, commands.Command) error),
	}

	appCommands.Register("login", commands.HandlerLogin)

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
