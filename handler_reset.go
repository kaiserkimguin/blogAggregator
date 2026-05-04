package main 

import(
	"fmt"
	"errors"
	"context"
	// "github.com/kaiserkimguin/blogAggregator/internal/database"
)

func handlerReset (s *state, cmd command) error {
	// check wether the right amount of args was passed
	if len(cmd.args) != 0 {
		return errors.New("unexpected argument")
	}
	// call Reset query function
	err := s.db.Reset(context.Background())
	if err != nil {
		fmt.Println("reset unsuccessful")
		return err
	}
	// print success message and return
	fmt.Println("Database was reset, all users dropped")
	return nil 
}
