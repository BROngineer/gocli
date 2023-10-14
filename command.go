package gocli

type Command struct {
	Name        string
	Subcommands map[string]Command
	Run         func(Command)
	RunE        func(Command) error
	FlagSet     FlagSet
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
	if c.Flags() == nil {
		c.FlagSet = NewFlagSet()
	}
	c.FlagSet.AddFlag(flag)
	return c
}

func (c Command) GetFlag(name string) Flag {
	return c.Flags()[name]
}

func (c Command) Execute() {
	c.Run(c)
}

func (c Command) ExecuteE() error {
	return c.RunE(c)
}
