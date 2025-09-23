package main

import (
	"os"

	"github.com/aman-zulfiqar/go-nihonvoice/cmd"
)

func main() {
	cli := &cmd.CLI{ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
