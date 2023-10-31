package gocli

import (
	"errors"
	"os"
	"path"
	"strings"
)

func Run(command *Command) error {
	var (
		err error
		cmd *Command
	)

	args := os.Args
	inheritFlags(command)
	cmd, err = evaluate(command, args)
	if err != nil {
		return err
	}
	err = validate(cmd)
	if err != nil {
		return err
	}
	return execute(cmd)
}

func inheritFlags(cmd *Command) {
	for _, subcommand := range cmd.Subcommands {
		for _, flag := range cmd.Flags {
			if flag.Shared() {
				subcommand.Flags[flag.Name()] = flag
			}
		}
		inheritFlags(&subcommand)
	}
}

func evaluate(cmd *Command, args []string) (*Command, error) {
	var (
		value string
		err   error
	)

	for i := 0; i < len(args); i++ {
		value = ""
		arg := args[i]
		if strings.HasPrefix(arg, "-") {
			arg = strings.TrimLeft(arg, "-")
			if strings.Contains(arg, "=") {
				arg, value = splitEqualsChar(arg)
			}
			flag := cmd.FlagByName(arg)
			if flag != nil {
				switch flag.Value().(type) {
				case *bool:
					value = "true"
					err = flag.parse(value)
				default:
					switch {
					case value == "" && i == len(args)-1:
						return nil, errors.New("error placeholder")
					case value == "" && strings.HasPrefix(args[i+1], "-"):
						return nil, errors.New("error placeholder")
					case value == "":
						value = args[i+1]
						i++
						fallthrough
					default:
						err = flag.parse(value)
					}
				}
				if err != nil {
					return nil, err
				}
				continue
			}
		}
		arg = sanitizeCommand(args[i])
		if arg == cmd.Name {
			continue
		}
		command, ok := cmd.Subcommands[arg]
		if ok {
			cmd, err = evaluate(&command, args[i:])
			if err != nil {
				return nil, err
			}
			return cmd, nil
		} else {
			return nil, errors.New("error placeholder")
		}
	}
	return cmd, nil
}

func validate(cmd *Command) error {
	for _, flag := range cmd.Flags {
		requiredNotParsed := flag.Required() && !flag.Parsed()
		optionalNoValue := !flag.Required() && flag.IsNilValue()
		if requiredNotParsed {
			return errors.New("required but not set")
		}
		if optionalNoValue {
			return errors.New("optional but no value")
		}
	}
	return nil
}

func execute(cmd *Command) error {
	if cmd.RunE != nil {
		return cmd.RunE(cmd)
	}
	if cmd.Run != nil {
		cmd.Run(cmd)
		return nil
	}
	return errors.New("no run func error placeholder")
}

func splitEqualsChar(in string) (string, string) {
	split := strings.Split(in, "=")
	return split[0], split[1]
}

func sanitizeCommand(in string) string {
	cleaned := path.Clean(in)
	file := path.Base(cleaned)
	return file
}
