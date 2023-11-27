package cmd

import (
	"gocli"
)

func RootCommand() *gocli.Command {
	return gocli.NewCommand("example",
		gocli.Flags(
			gocli.BoolFlag("verbose", gocli.Shorthand("v"), gocli.Shared()),
			gocli.StringFlag("log-level", gocli.Shorthand("l"), gocli.Shared(), gocli.Default[string]("info")),
			gocli.StringFlag("log-file", gocli.Shared(), gocli.Required()),
		),
		gocli.Subcommands(RunCommand()))
}
