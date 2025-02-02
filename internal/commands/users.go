package commands

import (
	"context"
	"fmt"

	"github.com/wolv89/bootgator/internal/state"
)

func HandlerUsers(s *state.State, cmd Command) error {

	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return err
	}

	var status string

	for _, user := range users {
		status = ""
		if user.Name == s.Config.CurrentUserName {
			status = " (current)"
		}
		fmt.Printf("* %s%s\n", user.Name, status)
	}

	return nil

}
