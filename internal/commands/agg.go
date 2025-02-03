package commands

import (
	"context"
	"fmt"

	"github.com/wolv89/bootgator/internal/rss"
	"github.com/wolv89/bootgator/internal/state"
)

func HandlerAgg(s *state.State, cmd Command) error {

	feed, err := rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println("##", feed.Channel.Title)
	fmt.Println(feed.Channel.Description)
	fmt.Println("---")

	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
		fmt.Println(item.Description[:125] + "...")
		fmt.Println("")
	}

	fmt.Println("")

	return nil

}
