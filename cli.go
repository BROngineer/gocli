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
	args := os.Args[1:]
	inheritFlags(&rootCommand)
	command, err := evaluate(rootCommand, args)
	if err != nil {
		return err
	}
	err = validateFlags(command.FlagSet)
	if err != nil {
		return err
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

func validateFlags(flags FlagSet) error {
	for _, f := range flags.Flags {
		if f.Parsed() {
			continue
		}
		if f.Required() {
			return fmt.Errorf("flag %s is required", f.Name())
		}
		if f.ValueOrDefault() == nil {
			return fmt.Errorf("optional flags require default value")
		}
	}
	return nil
}

// evaluate function will parse the os.Args and compare args slice
// with command passed as an argument
func evaluate(cmd Command, args []string) (Command, error) {
	var (
		err   error
		value string
	)

	for i := 0; i < len(args); i++ {
		value = ""
		arg := args[i]
		if arg == cmd.Name {
			continue
		}
		if strings.HasPrefix(arg, "-") {
			arg = strings.TrimLeft(arg, "-")
			if strings.Contains(arg, "=") {
				arg, value = splitEqualsChar(arg)
			}
			flag := cmd.Flag(arg)
			if flag != nil {
				switch flag.Value().(type) {
				case *bool:
					value = "true"
					err = flag.Parse(value)
				default:
					switch {
					case value == "" && i == len(args)-1:
						return Command{}, fmt.Errorf("no value passed for flag %s", flag.Name())
					case value == "" && strings.HasPrefix(args[i+1], "-"):
						return Command{}, fmt.Errorf("no value passed for flag %s", flag.Name())
					case value == "":
						value = args[i+1]
						i++
						fallthrough
					default:
						err = flag.Parse(value)
					}
				}
				if err != nil {
					return Command{}, err
				}
				flag.SetParsed()
				continue
			}
			return Command{}, fmt.Errorf("flag -%s is not defined for command %s", arg, cmd.Name)
		}
		command, ok := cmd.Subcommands[arg]
		if ok {
			cmd, err = evaluate(command, args[i:])
			if err != nil {
				return Command{}, err
			}
			return cmd, nil
		} else {
			return Command{}, fmt.Errorf("unrecognized command %s", arg)
		}
	}
	return cmd, nil
}

func splitEqualsChar(in string) (string, string) {
	split := strings.Split(in, "=")
	return split[0], split[1]
}
