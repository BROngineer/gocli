package gocli

import (
	"fmt"
	"os"
	"strings"
)

type CLIApp struct {
	command   Command
	evaluated bool
}

func NewCLIApp() *CLIApp {
	return &CLIApp{}
}

func (c *CLIApp) Evaluate(rootCommand Command) error {
	args := os.Args
	inheritFlags(&rootCommand)
	command, err := evaluate(rootCommand, args)
	if err != nil {
		return err
	}
	if !validateRequiredFlags(command.FlagSet) {
		return fmt.Errorf("not all required flags passed")
	}
	c.command = command
	c.evaluated = true
	return nil
}

func (c *CLIApp) Execute() error {
	if c.evaluated {
		return c.command.Execute()
	}
	return fmt.Errorf("input have to be evaluated before execution")
}

func inheritFlags(cmd *Command) {
	for _, subcommand := range cmd.Subcommands {
		for _, flag := range cmd.Flags() {
			if flag.Shared() {
				subcommand.WithFlag(flag)
			}
		}
		inheritFlags(&subcommand)
	}
}

func validateRequiredFlags(flags FlagSet) bool {
	for _, f := range flags.Flags {
		if f.Parsed() {
			continue
		}
		if f.Required() {
			return false
		}
	}
	return true
}

// evaluate function will parse the os.Args and compare args slice
// with command passed as an argument
func evaluate(cmd Command, args []string) (Command, error) {
	var err error
	var value string
	for i := 0; i < len(args); i++ {
		value = ""
		item := args[i]
		if item == cmd.Name {
			continue
		}
		item = strings.TrimLeft(item, "-")
		if strings.Contains(item, "=") {
			item, value = splitEqualsChar(item)
		}
		flag := cmd.Flag(item)
		if flag != nil {
			switch flag.Value().(type) {
			case *bool:
				value = "true"
				err = flag.Parse(value)
			default:
				if value == "" {
					if i == len(args)-1 {
						return Command{}, fmt.Errorf("no value provided")
					}
					value = args[i+1]
					i++
				}
				err = flag.Parse(value)
			}
			if err != nil {
				return Command{}, err
			}
			flag.SetParsed()
			continue
		}
		command, ok := cmd.Subcommands[item]
		if ok {
			cmd, err = evaluate(command, args[i:])
			if err != nil {
				return Command{}, err
			}
			return cmd, nil
		}
	}
	return cmd, nil
}

func splitEqualsChar(in string) (string, string) {
	split := strings.Split(in, "=")
	return split[0], split[1]
}
