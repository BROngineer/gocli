package gocli

import (
	"errors"
)

type CommandsSet map[string]Command
type FlagsSet map[string]Flag

type Command struct {
	Name        string
	Description string
	Flags       FlagsSet
	Subcommands CommandsSet
	Run         func(*Command)
	RunE        func(*Command) error
}

func NewCommand(name string, opts ...CommandOption) *Command {
	cmd := &Command{
		Name:        name,
		Flags:       make(FlagsSet),
		Subcommands: make(CommandsSet),
	}
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd
}

func (c *Command) FlagByName(name string) Flag {
	flag, ok := c.Flags[name]
	if ok {
		return flag
	}
	for _, item := range c.Flags {
		if name == item.Shorthand() {
			return item
		}
	}
	return flag
}

func FlagValue[T allowed](cmd *Command, flagName string) (*T, error) {
	flag := cmd.FlagByName(flagName)
	if flag == nil {
		return nil, errors.New("flag not found")
	}
	if flag.IsNilValue() {
		return nil, errors.New("flag value is not defined")
	}
	res, ok := flag.Value().(*T)
	if ok {
		return res, nil
	}
	return nil, errors.New("wrong type")
}
