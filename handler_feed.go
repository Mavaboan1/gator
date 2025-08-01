package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"main/internal/database"
	"time"
)

func handlerFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	userId := user.ID
	name := cmd.Args[0]
	url := cmd.Args[1]

	createdFeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    userId,
	})
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    userId,
		FeedID:    createdFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed created successfully:")
	printFeed(createdFeed, user)
	fmt.Println()
	fmt.Println("Feed followed successfully:")
	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	fmt.Println("=====================================")
	return nil

}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
	fmt.Printf("* User:          %s\n", user.Name)
}
