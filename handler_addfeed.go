package main

import(
	"errors"
	"context"
	"time"
	"fmt"
	"github.com/google/uuid"
	"github.com/kaiserkimguin/blogAggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error{
	// check for correct no of args
	if len (cmd.args) != 2 {
		return errors.New("two arguments expected")
	}
	// get current user to connect to feed
	curUserName := s.cfg.CurrentUserName
	curUser, err := s.db.GetUser(context.Background(), curUserName)
	if err != nil {
		return err
	}
	// Create empty feed to fil in following data.
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:					uuid.New(),
		CreatedAt:	time.Now(),
		UpdatedAt:	time.Now(),
		Name: 			cmd.args[0],
		Url:				cmd.args[1],
		UserID:			curUser.ID,
	})
	if err != nil {
		fmt.Println("unable to create feed")
		return err
	}
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
