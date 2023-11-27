package main

import (
	"fmt"
	"os"

	"gocli"
	"gocli/example/cmd"
)

func main() {
	err := gocli.Run(cmd.RootCommand())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
