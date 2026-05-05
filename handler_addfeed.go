package main

import(
	"errors"
	"context"
	"time"
	"fmt"
	"github.com/google/uuid"
	"github.com/kaiserkimguin/blogAggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error{
	// check for correct no of args
	if len (cmd.args) != 2 {
		return errors.New("two arguments expected")
	}
		// Create feed with the given data 
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:					uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name: 			cmd.args[0],
		Url:				cmd.args[1],
		UserID:			user.ID,
	})
	if err != nil {
		fmt.Println("unable to create feed")
		return err
	}
	// create new feedfollow with constructed feed
	_, err = s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID: 				uuid.New(),
			CreatedAt:	time.Now(),
			UpdatedAt: 	time.Now(),
			UserID:			user.ID,
			FeedID: 		feed.ID,
		})
	// print the recorded feed to console
	fmt.Printf("%+v\n", feed)
	return nil
}

func handlerFeeds(s* state, cmd command) error{
	// no args check
	if len(cmd.args) != 0 {
		return errors.New("no arguments expected")
	}
	// get all feeds from db using getFeeds method
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err 
	}
	// print all feeds to console
	for index, feed := range feeds {
		fmt.Printf("%d: %+v\n", index, feed)
	}
	return nil
}
