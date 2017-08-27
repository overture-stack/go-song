package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "Upload analysis metadata for validation",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "Check validation status of an uploaded analysis",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	app.Run(os.Args)
}
