package commands

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerFollowing(s *state.State, cmd Command) error {

	if len(s.Config.CurrentUserName) == 0 {
		return fmt.Errorf("please login to follow a feed")
	}

	foundUser, _ := s.DB.GetUser(context.Background(), s.Config.CurrentUserName)
	if foundUser.Name != s.Config.CurrentUserName {
		return fmt.Errorf("user: %s does not exist", s.Config.CurrentUserName)
	}

	myFeeds, err := s.DB.GetFeedFollowsForUser(context.Background(), uuid.NullUUID{
		UUID:  foundUser.ID,
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
