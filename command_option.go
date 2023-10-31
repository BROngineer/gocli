package cli

type CommandOption func(cmd *Command)

func CommandDescription(value string) CommandOption {
	return func(cmd *Command) {
		cmd.Description = value
	}
}

func Subcommands(commands ...*Command) CommandOption {
	return func(cmd *Command) {
		for _, c := range commands {
			cmd.Subcommands[c.Name] = *c
		}
	}
}

func RunFunction(f func(*Command)) CommandOption {
	return func(cmd *Command) {
		cmd.Run = f
	}
}

func RunEFunction(f func(*Command) error) CommandOption {
	return func(cmd *Command) {
		cmd.RunE = f
	}
}

func Flags(flags ...Flag) CommandOption {
	return func(cmd *Command) {
		for _, flag := range flags {
			cmd.Flags[flag.Name()] = flag
		}
	}
}
