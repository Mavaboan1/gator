package main

import (
	"context"
	"fmt"
	"main/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("unfollow expects a url to unfollow")
	}
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.Args[0])
	if err != nil {
		return err
	}

	err = s.db.DeleteFeed(context.Background(), database.DeleteFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Unfollowed feed: %s\n", feed.Name)
	fmt.Printf("Unfollowed url: %s\n", feed.Url)
	return nil
}
