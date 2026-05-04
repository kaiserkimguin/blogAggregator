package main

import (
	"fmt"
	"errors"
	"context"
)

func handlerAgg(s *state, cmd command) error{
	// check wether a url argument was passed
	if len(cmd.args) != 0 {
		return errors.New(" no arguments expected")
	}
	// hardcoding testUrl for testing purposes
	testUrl := "https://www.wagslane.dev/index.xml"
	// call fetch feed with url and context to obtain RSSFeed
	RF, err := fetchFeed(context.Background(), testUrl) // should be cmd.args[0]
	if err != nil {
		return err
	}
	// print RSSFeed to console
	fmt.Printf("%+v\n", RF)
	// return non-error
	return nil 
}
