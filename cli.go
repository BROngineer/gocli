package gocli

import (
	"fmt"
	"strings"
)

func Run(cmd Command, args []string) error {
	inheritFlags(&cmd)
	command, err := evaluate(cmd, args)
	if err != nil {
		return err
	}
	if command.RunE != nil {
		return command.ExecuteE()
	}
	if command.Run != nil {
		command.Execute()
		return nil
	}
	return fmt.Errorf("no function to run")
}

func inheritFlags(cmd *Command) {
	for _, subcommand := range cmd.Subcommands {
		for _, flag := range cmd.Flags() {
			if !flag.Shared() {
				continue
			}
			subcommand.FlagSet.AddFlag(flag)
		}
		inheritFlags(&subcommand)
	}
}

// evaluate function will parse the os.Args and compare args slice
// with command passed as an argument
func evaluate(cmd Command, args []string) (Command, error) {
	flagSet := NewFlagSet()
	var err error
	for i := 0; i < len(args); i++ {
		item := args[i]
		if item == cmd.Name {
			continue
		}
		item = strings.TrimLeft(item, "-")
		flag, ok := cmd.Flags()[item]
		if ok {
			switch flag.Value().(type) {
			case *bool:
				err = flag.Parse("")
			default:
				err = flag.Parse(args[i+1])
			}
			if err != nil {
				return Command{}, err
			}
			if flag.Shared() {
				flagSet.AddFlag(flag)
			}
			i++
			continue
		}
		// lookup for subcommands
		command, ok := cmd.Subcommands[item]
		if ok {
			cmd, err = evaluate(command, args[i:])
			if err != nil {
				return Command{}, err
			}
			cmd.FlagSet.Merge(&flagSet)
			return cmd, nil
		}
	}
	return cmd, nil
}
