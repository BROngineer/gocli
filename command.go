package gocli

type Command struct {
	Name        string
	Subcommands map[string]Command
	Run         func(Command)
	RunE        func(Command) error
	FlagSet     FlagSet
	Config      Config
}

func NewCommand(name string) Command {
	return Command{
		Name:    name,
		FlagSet: NewFlagSet(),
	}
}

func (c Command) Flags() map[string]Flag {
	return c.FlagSet.Flags
}

func (c Command) Flag(name string) Flag {
	return c.FlagSet.Flag(name)
}

func (c Command) WithSubcommand(cmd Command) Command {
	if c.Subcommands == nil {
		c.Subcommands = make(map[string]Command)
	}
	c.Subcommands[cmd.Name] = cmd
	return c
}

func (c Command) WithRunFunc(f func(Command)) Command {
	c.Run = f
	return c
}

func (c Command) WithRunEFunc(f func(Command) error) Command {
	c.RunE = f
	return c
}

func (c Command) WithFlag(flag Flag) Command {
	c.FlagSet.AddFlag(flag)
	return c
}

func (c Command) WithConfig(cfg Config) Command {
	c.Config = cfg
	return c
}

func (c Command) Execute() error {
	if c.RunE != nil {
		return c.RunE(c)
	}
	if c.Run != nil {
		c.Run(c)
		return nil
	}
	return CommandExecuteError()
}
