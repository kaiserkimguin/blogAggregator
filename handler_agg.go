package main

import (
	"fmt"
	"errors"
	"context"
	"html"
	"time"
	//"github.com/kaiserkimguin/blogAggregator/internal/database"
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
	fmt.Println(html.UnescapeString(RF.Channel.Title))
	for index, item := range RF.Channel.Item{
		fmt.Println(index,": ",html.UnescapeString(item.Title))
	}
	return nil
} 
