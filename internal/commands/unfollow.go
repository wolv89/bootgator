package commands

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/wolv89/bootgator/internal/database"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerUnfollow(s *state.State, cmd Command, user database.User) error {

	if len(cmd.Args) < 1 {
		return fmt.Errorf("need a url of feed to unfollow")
	}

	feedurl := cmd.Args[0]

	foundFollow, err := s.DB.GetFeedFollowForUserFromURL(context.Background(), database.GetFeedFollowForUserFromURLParams{
		UserID: uuid.NullUUID{
			UUID:  user.ID,
			Valid: true,
		},
		Url: feedurl,
	})
	if err != nil {
		return err
	}

	err = s.DB.DeleteFeedFollow(context.Background(), foundFollow.ID)
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("Stopped following feed:", foundFollow.FeedName)
	fmt.Println("")

	return nil

}
