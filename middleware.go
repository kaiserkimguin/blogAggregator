package main

import(
	"context"
	"github.com/kaiserkimguin/blogAggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error{
	return func(s *state, cmd command) error {
		curUserName := s.cfg.CurrentUserName
		curUser, err := s.db.GetUser(context.Background(), curUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, curUser)
	}
}
