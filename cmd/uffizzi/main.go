package main

import (
	"fmt"
	"log"
	"os"

	"UffizziCloud/uffizzi-cli/internal/commands/config"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    "uffizzi",
		Usage:   "A command-line interace (CLI) for Uffizzi App",
		Suggest: true,
	}
	app.Commands = []*cli.Command{
		{
			Name:    "config",
			Aliases: []string{"t"},
			Usage: `The uffizzi config command lets you configure this command-line application.
				If COMMAND is not specified, uffizzi config launches an interactive set up
				guide.
				
				For more information on configuration options, see:
				https://github.com/UffizziCloud/uffizzi_cli`,
			Action: func(c *cli.Context) error {
				config.Config()
				return nil
			},
			Subcommands: []*cli.Command{
				{
					Name:  "list",
					Usage: "Lists all options and their values from the config file.",
					Action: func(cCtx *cli.Context) error {
						config.List()
						return nil
					},
				},
				{
					Name:  "get-value",
					Usage: "Displays the value of the specified option.",
					Action: func(cCtx *cli.Context) error {
						config.GetValue(cCtx.Args().First())
						return nil
					},
				},
				{
					Name:  "set",
					Usage: "Sets the value of the specified option.",
					Action: func(cCtx *cli.Context) error {
						config.SetValue(cCtx.Args().Get(0), cCtx.Args().Get(1))
						return nil
					},
				},
				{
					Name:  "unset",
					Usage: "Deletes the value of the specified option.",
					Action: func(cCtx *cli.Context) error {
						fmt.Println("unsetset command called: ", cCtx.Args().First())
						return nil
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
