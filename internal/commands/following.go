package commands

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerFollowing(s *state.State, cmd Command, user database.User) error {

	myFeeds, err := s.DB.GetFeedFollowsForUser(context.Background(), uuid.NullUUID{
		UUID:  user.ID,
		Valid: true,
	})
	if err != nil {
		return err
	}

	if len(myFeeds) == 0 {
		fmt.Println("whoops, you're not following anything...")
		return nil
	}

	fmt.Println("")
	fmt.Println("Your Feeds")
	fmt.Println("----------")
	for _, feed := range myFeeds {
		fmt.Println(feed)
	}
	fmt.Println("")

	return nil

}
