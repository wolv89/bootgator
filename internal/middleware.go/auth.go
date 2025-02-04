package middleware

import (
	"context"
	"fmt"

	"github.com/wolv89/bootgator/internal/commands"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/state"
)

func LoggedIn(handler func(s *state.State, cmd commands.Command, user database.User) error) func(*state.State, commands.Command) error {

	return func(s *state.State, cmd commands.Command) error {

		if len(s.Config.CurrentUserName) == 0 {
			return fmt.Errorf("please login to add a feed")
		}

		foundUser, _ := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
		if foundUser.Name != s.Config.CurrentUserName {
			return fmt.Errorf("user: %s does not exist", s.Config.CurrentUserName)
		}

		return handler(s, cmd, foundUser)

	}

}
