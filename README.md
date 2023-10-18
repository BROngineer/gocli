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

### Usage

```go
package main

import (
  "fmt"
  "os"

  "gocli"
)

func runCommandExec(cmd gocli.Command) {
  env, _ := gocli.GetValue[string](cmd.FlagSet, "env")
  verbose, _ := gocli.GetValue[bool](cmd.FlagSet, "verbose")
  fmt.Printf("Running in %s environment\n", *env)
  fmt.Printf("Verbose output: %v\n", *verbose)
}

var (
  runCommand = gocli.NewCommand("run").
    WithFlag(gocli.NewFlag[string]("env", "").WithDefault("DEV")).
    WithRunFunc(runCommandExec)

  rootCommand = gocli.NewCommand("example").
    WithSubcommand(runCommand).
    WithFlag(gocli.NewFlag[bool]("verbose", "").SetShared())
)

func main() {
  err := gocli.Run(rootCommand, os.Args)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  os.Exit(0)
}
```
```shell
> example run
# Output
Running in DEV environment
Verbose output: false

> example run -env=prod -verbose
# Output
Running in prod environment
Verbose output: true
```