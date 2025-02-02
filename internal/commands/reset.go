package commands

import (
	"context"
	"fmt"

	"github.com/wolv89/bootgator/internal/state"
)

func HandlerReset(s *state.State, cmd Command) error {

	err := s.DB.Reset(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("Resetting users table")

	return nil

}
