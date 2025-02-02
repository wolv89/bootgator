package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerRegister(s *state.State, cmd Command) error {

	if len(cmd.Args) == 0 {
		return fmt.Errorf("register command expects a username argument")
	}

	username := cmd.Args[0]

	foundUser, _ := s.DB.GetUser(context.Background(), username)
	if foundUser.Name == username {
		return fmt.Errorf("user: %s already exists", username)
	}

	now := time.Now()

	newUser, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	})
	if err != nil {
		return err
	}

	fmt.Println("User created:", username)
	fmt.Println(newUser)

	s.Config.SetUser(username)
	fmt.Println("Setting current user to:", username)

	return nil

}
