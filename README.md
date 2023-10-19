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

##### Command custom config example

Here is an example of how to add and leverage  user-defined config structure for `gocli.Command`

```go
package main

import (
  "fmt"
  "os"

  "github.com/brongineer/gocli"
)

type CmdConfig struct {
  Message string
  Verbose bool
}

func (c *CmdConfig) LoadFromFlags(flags gocli.FlagSet) error {
  // your custom logic here, for instance:
  m, _ := gocli.GetValue[string](flags, "message")
  c.Message = *m
  v, _ := gocli.GetValue[bool](flags, "verbose")
  c.Verbose = *v
  return nil
}

func (c *CmdConfig) LoadFromEnv(prefix string) error {
  // your custom logic here
  return nil
}

func (c *CmdConfig) LoadFromFiles(files []string) error {
  // your custom logic here
  return nil
}

func actualWork(cfg *CmdConfig) error {
  if cfg.Verbose {
    fmt.Printf("This is verbose message: %s\n", cfg.Message)
	return nil
  }
  fmt.Println(cfg.Message)
  return nil
}

func runCommand(cmd gocli.Command) error {
  err := cmd.Config.LoadFromFlags(cmd.FlagSet)
  if err != nil {
    return err
  }
  cfg := gocli.TypedConfig[CmdConfig](cmd.Config)
  return actualWork(cfg)
}

var command = gocli.NewCommand("example").
  WithFlag(gocli.NewFlag[bool]("verbose", "")).
  WithFlag(gocli.NewFlag[string]("message", "")).
  WithRunEFunc(runCommand).
  WithConfig(&CmdConfig{})

func main() {
  _ = gocli.Run(command, os.Args)
}
```
```shell
> example -message=hello
# Output
hello

> example -message=hello -verbose
# Output
This is verbose message: hello
```