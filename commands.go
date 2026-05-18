package main

import "fmt"

type Command struct {
	name string
	args []string
}

type Commands struct {
	cmds map[string]func(*State, Command) error
}

func (c *Commands) run(s *State, cmd Command) error {
	if s == nil {
		return fmt.Errorf("state is not exists")
	}

	handler, exists := c.cmds[cmd.name]

	if !exists {
		return fmt.Errorf("this handler not exists")
	}

	if err := handler(s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.cmds[name] = f
}
