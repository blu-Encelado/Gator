package main

import (
	"Gator/internal/config"
	"Gator/internal/database"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return fmt.Errorf("not existing state")
	}

	err := c.registeredCommands[cmd.name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("command arguments empty")
	}
	_, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("user don't exist: %w", err)
	}

	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println("user set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("command arguments empty")
	}
	context_var := context.Background()
	uuid_var := uuid.New()
	time_now := time.Now()
	user_new := database.CreateUserParams{
		ID:        int32(uuid_var.ID()),
		CreatedAt: time_now,
		UpdatedAt: time_now,
		Name:      cmd.arguments[0]}

	user, err := s.db.CreateUser(context_var, user_new)
	if err != nil {
		return fmt.Errorf("fail to generate User: %w", err)
	}
	err = s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println("User created:")
	fmt.Println(user)

	return nil
}

func handlerReset(s *state, cmd command) error {
	fmt.Printf("Called: %s\n", cmd.name)

	context_var := context.Background()
	err := s.db.Reset(context_var)
	if err != nil {
		os.Exit(1)
		return fmt.Errorf("fail to reset: %w", err)
	}

	return nil
}
