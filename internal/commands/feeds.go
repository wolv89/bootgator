package commands

import (
	"context"
	"fmt"

	"github.com/wolv89/bootgator/internal/state"
)

func HandlerFeeds(s *state.State, cmd Command) error {

	feeds, err := s.DB.GetFeedsWithUsernames(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.Username)
		fmt.Println("")
	}

	return nil

}
