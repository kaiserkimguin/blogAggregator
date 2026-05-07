package main

import (
	"database/sql"
	"fmt"
	"errors"
	"context"
	"html"
	"time"
	"github.com/google/uuid"
	"github.com/kaiserkimguin/blogAggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error{
	// check wether a url argument was passed
	if len(cmd.args) != 1 {
		return errors.New("one arguments expected")
	}
	// parse the argument into time.duration format
	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	// Print duration to terminal and use it to make requests using a ticker
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)	
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	return nil 
}

func scrapeFeeds(s *state) error{
	// call next feed to fetch, no args needed
	nextFeed, err :=	s.db.GetNextFeedToFetch(context.Background())	
	if err != nil{
		return err
	}
	fmt.Printf("%s is the next feed to fetch\n", nextFeed.Name)
	//mark feed as fetched
	markedFeed, err := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return err
	}
	// fetch the feed and print its titles to console
	RF, err := fetchFeed(context.Background(), markedFeed.Url)
	if err != nil {
		return err
	}
	// save the posts to the posts table
	fmt.Println(html.UnescapeString(RF.Channel.Title))
	for _, item := range RF.Channel.Item{
		// parse the time
		parsedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err // try another time format if this fails
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:						uuid.New(), CreatedAt:		time.Now(),
			UpdatedAt:		time.Now(),
			Title:				item.Title,
			Url:					item.Link,	
			Description: 	sql.NullString{
    		String: item.Description,
    		Valid:  item.Description != "",
				},
			PublishedAt:	sql.NullTime{
			 Time:  	parsedTime,
		   Valid:	  err == nil,
		 		},
			FeedID:				markedFeed.ID,
		})
		if err != nil {
			return err
		}
	}
	return nil
} 
