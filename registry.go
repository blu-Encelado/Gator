package main

import (
	"Gator/internal/config"
	"fmt"
)

type state struct {
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
	err := s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println("user set")
	return nil
}
