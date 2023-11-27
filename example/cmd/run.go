package cmd

import (
	"fmt"

	"gocli"
)

func DoWork(cmd *gocli.Command) {
	verbose, _ := gocli.FlagValue[bool](cmd, "verbose")
	logLevel, _ := gocli.FlagValue[string](cmd, "log-level")
	logFile, _ := gocli.FlagValue[string](cmd, "log-file")
	host, _ := gocli.FlagValue[string](cmd, "h")
	port, _ := gocli.FlagValue[int](cmd, "p")
	if *verbose {
		fmt.Printf("running on %s:%d; log level: %s, log file: %s\n", *host, *port, *logLevel, *logFile)
		return
	}
	fmt.Println(*host, *port, *logLevel, *logFile)
}

func RunCommand() *gocli.Command {
	return gocli.NewCommand("run",
		gocli.Flags(
			gocli.StringFlag("host", gocli.Shorthand("h"), gocli.Default[string]("127.0.0.1")),
			gocli.IntFlag("port", gocli.Shorthand("p"), gocli.Default[int](80)),
		),
		gocli.RunFunction(DoWork),
	)
}
