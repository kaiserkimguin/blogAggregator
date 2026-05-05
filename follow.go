package main 

import (
	"fmt"
	"context"
	"time"
	"errors"
	"github.com/google/uuid"
	"github.com/kaiserkimguin/blogAggregator/internal/database"
)


func handlerFollow(s *state, cmd command, user database.User) error{
	// check for correct number of arguments
	if len(cmd.args) != 1 {
		return errors.New("Exactly one argument (url) expected")
	}
		// look up feed by url using GetFeed
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	// insert feed using CreateFeedFollow
	feedFol, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:					uuid.New(),
			CreatedAt: 	time.Now(),
			UpdatedAt: 	time.Now(),
			UserID:			user.ID,
			FeedID:			feed.ID,
		})
	if err !=nil {
		return err
	}
	// print feedFol, which includes names of feed and user, return nil
	fmt.Printf("%+v\n", feedFol)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error{
	if len(cmd.args) != 0 {
		return errors.New("no arguments expected")
	}
		// use query to obtain requested data 
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	// print the names of the feeds and return nil
	for index, feed := range feeds{
		fmt.Printf("%d: %+v\n", index, feed.Feedname)
	}
	return nil
}
