package gocli

import (
	"fmt"
	"path"
	"strings"
)

func Run(rootCommand Command, args []string) error {
	var (
		cmd Command
		err error
	)

	inheritFlags(&rootCommand)
	cmd, err = evaluate(rootCommand, args)
	if err != nil {
		return err
	}
	err = validateFlags(cmd.FlagSet)
	if err != nil {
		return err
	}
	return cmd.Execute()
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
			return NewFlagError(fmt.Sprintf(MissedRequiredFlagErrorMessage, f.Name()))
		}
		if f.ValueOrDefault().IsNil() {
			return NewFlagError(fmt.Sprintf(MissedDefaultValueErrorMessage, f.Name()))
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
		if strings.HasPrefix(arg, "-") {
			arg = strings.TrimLeft(arg, "-")
			if strings.Contains(arg, "=") {
				arg, value = splitEqualsChar(arg)
			}
			flag := cmd.Flag(arg)
			if flag != nil {
				switch flag.Value().Value().(type) {
				case *bool:
					value = "true"
					err = flag.Parse(value)
				default:
					switch {
					case value == "" && i == len(args)-1:
						return Command{}, NewFlagError(fmt.Sprintf(MissedFlagValueErrorMessage, flag.Name()))
					case value == "" && strings.HasPrefix(args[i+1], "-"):
						return Command{}, NewFlagError(fmt.Sprintf(MissedFlagValueErrorMessage, flag.Name()))
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
			return Command{}, NewCommandError(fmt.Sprintf(UndefinedFlagErrorMessage, arg, cmd.Name))
		}
		arg = sanitizeCommand(args[i])
		if arg == cmd.Name {
			continue
		}
		command, ok := cmd.Subcommands[arg]
		if ok {
			cmd, err = evaluate(command, args[i:])
			if err != nil {
				return Command{}, err
			}
			return cmd, nil
		} else {
			return Command{}, NewCommandError(fmt.Sprintf(UndefinedCommandErrorMessage, arg))
		}
	}
	return cmd, nil
}

func splitEqualsChar(in string) (string, string) {
	split := strings.Split(in, "=")
	return split[0], split[1]
}

func sanitizeCommand(input string) string {
	cleaned := path.Clean(input)
	file := path.Base(cleaned)
	return file
}
