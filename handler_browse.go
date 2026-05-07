package main

import(
	"errors"
	"context"
	"fmt"
	"strconv"
	"github.com/kaiserkimguin/blogAggregator/internal/database"

)

func handlerBrowse(s *state, cmd command, user database.User) error{
	// check for args 
	if len(cmd.args) > 1 {
		return errors.New("0 or 1(number of posts) arguments expected")
	}
	// set default limit to 2, if not chosen by user
	lim := 2
	if len(cmd.args) == 1 {
    if specifiedLimit, err := strconv.Atoi(cmd.args[0]); err == nil {
        lim = specifiedLimit
    } else {
        return fmt.Errorf("invalid limit: %w", err)
    }
	}
	//get Posts for User 
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		 UserID: 	user.ID,
		 Limit:		int32(lim),
	})
	if err != nil {
		return err
	}
	// print posts to terminal
	for index, post := range posts {
		fmt.Printf("%d: %+v\n", index, post)
	}	
	return nil
}
