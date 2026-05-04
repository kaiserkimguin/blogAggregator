package main 

import (
	"errors"
	"fmt"
	"context"
	"time"
	"os"
	"github.com/kaiserkimguin/blogAggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
 // check wether the correct number of args are provided
 if len(cmd.args) != 1 {
	 return errors.New("login expects exactly 1 argument")
 }
 // check wether the user is already in the database
 _, err := s.db.GetUser(context.Background(), cmd.args[0]) 
 if err != nil {
	 fmt.Println("Login refused. User not registered", err)
	 os.Exit(1)
 }
 // to login, set current user in config file
 // use state to access the file
 // set username to provided arguments
 err = s.cfg.SetUser(cmd.args[0])
 if err != nil {
	 return err
 }
 // print a success msg 
 fmt.Printf("current user is now: %s", cmd.args[0])
 return nil
}

func handlerRegister(s *state, cmd command) error {
 // check wether the correct number of args are provided
 if len(cmd.args) != 1 {
	 return errors.New("register expects exactly 1 argument")
 }
 // create an empty query, for the following user data 
 user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
    ID:        uuid.New(),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
    Name:      cmd.args[0],
})
if err != nil {
	fmt.Println("unable to create user", err)
	os.Exit(1)
}
// print the created user to console for debugging
fmt.Printf("%+v\n", user) 
// log in the created user
if err = s.cfg.SetUser(cmd.args[0]); err != nil {
	return err
}
// Print success message and return absence of error.
fmt.Println("user successfully registered")
return nil
}

func handlerUsers(s *state, cmd command) error {
 // check wether the correct number of args are provided
 if len(cmd.args) != 0 {
	 return errors.New("no arguments expected")
 }
 // get all Users using states db.GetUsers method
 users, err := s.db.GetUsers(context.Background())
 if err != nil {
	 return err
 }
 // get current User from gatorconfig.
 curUser := s.cfg.CurrentUserName
 // print all users. Mark the current user
 for _, userX := range users {
	 if userX == curUser {
		 fmt.Println("* "+ userX + " (current)")
	 } else {
		 fmt.Println("* " + userX)
	 }
 }
 // return no error
 return nil 
}
