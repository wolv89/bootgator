package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerAddFeed(s *state.State, cmd Command, user database.User) error {

	if len(cmd.Args) < 2 {
		return fmt.Errorf("need a name and url to add a new feed")
	}

	feedname, feedurl := cmd.Args[0], cmd.Args[1]

	foundFeed, _ := s.DB.GetFeed(context.Background(), feedname)
	if foundFeed.Name == feedname {
		return fmt.Errorf("feed: %s already exists", feedname)
	}

	now := time.Now()

	newFeed, err := s.DB.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      feedname,
		Url:       feedurl,
		UserID: uuid.NullUUID{
			UUID:  user.ID,
			Valid: true,
		},
	})

	if err != nil {
		return err
	}

	now = time.Now()

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID: uuid.NullUUID{
			UUID:  user.ID,
			Valid: true,
		},
		FeedID: uuid.NullUUID{
			UUID:  newFeed.ID,
			Valid: true,
		},
	})

	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Saved new feed:", newFeed.Name)
	fmt.Println(newFeed)
	fmt.Println("")

	return nil

}
