package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerFollow(s *state.State, cmd Command, user database.User) error {

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
			UUID:  user.ID,
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
