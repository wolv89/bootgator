package commands

import (
	"fmt"

	"github.com/wolv89/bootgator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command expects a username argument")
	}

	s.Config.SetUser(cmd.Args[0])
	fmt.Println("Settings current user to:", cmd.Args[0])

	return nil

}
