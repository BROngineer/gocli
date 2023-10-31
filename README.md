# gocli

Minimalistic library to build simple CLI application with support of:

- Sub-commands
- Mandatory flags
- Shared flags: will be inherited by sub-commands
- Flag types:
  1. string
  2. int
  3. bool
  4. []string
  5. time.Duration
- Shorthands for flags 

### Usage

##### Subcommand example

```go
package main

import (
  "fmt"
  "os"

  "github.com/brongineer/gocli"
)

// root command where CLI app execution starts from
var rootCommand = gocli.NewCommand("example",
    gocli.Flags(
      gocli.BoolFlag("verbose", gocli.Shorthand("v"), gocli.Shared()),
      gocli.StringFlag("log-level", gocli.Shorthand("l"), gocli.Shared(), gocli.Default[string]("info")),
      gocli.StringFlag("log-file", gocli.Shared(), gocli.Required()),
    ),
    gocli.Subcommands(runCommand))

// subcommand which contains function doing actual work
var runCommand = gocli.NewCommand("run", 
    gocli.Flags(
      gocli.StringFlag("host", gocli.Shorthand("h"), gocli.Default[string]("127.0.0.1")), 
      gocli.IntFlag("port", gocli.Shorthand("p"), gocli.Default[int](80)), 
    ), 
    gocli.RunFunction(DoWork))

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

func main() {
  err := gocli.Run(rootCommand)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  os.Exit(0)
}
```
```shell
# flags "--log-level", "--log-file" and "--verbose" (shorthand "-v") can be placed 
# either before or after "run" subcommand since they are declared as shared and will 
# be inherited by all subcommand of root command

> example run --log-level trace -log-file /var/log/example.log -v
# Output
running on 127.0.0.1:80; log level: trace, log file: /var/log/example.log

> example run --log-level trace -log-file /var/log/example.log -p 8081
# Output
127.0.0.1 8081 trace /var/log/example.log
```
