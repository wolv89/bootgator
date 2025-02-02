package commands

import (
	"context"
	"fmt"

	"github.com/wolv89/bootgator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("login command expects a username argument")
	}

	username := cmd.Args[0]

	foundUser, _ := s.DB.GetUser(context.Background(), username)
	if foundUser.Name != username {
		return fmt.Errorf("user: %s does not exist", username)
	}

	s.Config.SetUser(username)
	fmt.Println("Setting current user to:", username)

	return nil

}
