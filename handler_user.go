package main 

import (
	"errors"
	"fmt"
)

func handlerLogin (s *state, cmd command) error {
 // check wether the correct number of args are provided
 if len(cmd.args) != 1 {
	 return errors.New("login expects exactly 1 argument")
 }
 // to login, set current user in config file
 // use state to access the file
 // set username to provided arguments
 err := s.cfg.SetUser(cmd.args[0])
 if err != nil {
	 return err
 }
 // print a success msg 
 fmt.Printf("current user is now: %s", cmd.args[0])
 return nil
}
