package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerFollow(s *state.State, cmd Command) error {

	if len(s.Config.CurrentUserName) == 0 {
		return fmt.Errorf("please login to follow a feed")
	}

	foundUser, _ := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if foundUser.Name != s.Config.CurrentUserName {
		return fmt.Errorf("user: %s does not exist", s.Config.CurrentUserName)
	}

	if len(cmd.Args) < 1 {
		return fmt.Errorf("need a url to follow feed")
	}

	feedurl := cmd.Args[0]

	foundFeed, _ := s.DB.GetFeedFromURL(context.Background(), feedurl)
	if foundFeed.Url != feedurl {
		return fmt.Errorf("feed: %s not found", feedurl)
	}

	now := time.Now()

	newFeed, err := s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: uuid.NullUUID{
			UUID:  foundUser.ID,
			Valid: true,
		},
		FeedID: uuid.NullUUID{
			UUID:  foundFeed.ID,
			Valid: true,
		},
	})

	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Now following feed:", newFeed.FeedName)
	fmt.Println(newFeed)
	fmt.Println("")

	return nil

}
