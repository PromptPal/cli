package main

import (
	"os"

	"github.com/PromptPal/cli/commands"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
)

var GitCommit string

func main() {
	app := &cli.App{
		Name:        "PromptPal Cli",
		Description: "PromptPal command line interface for generate types and more",
		Usage:       "PromptPal command line interface for generate types and more",
		Version:     GitCommit,
		Commands: []*cli.Command{
			commands.InitCommand,
			commands.GenerateCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red.Println(err)
	}
}
