package main

import (
	"gopkg.in/urfave/cli.v1"
)

// GetCommands returns a list of available commands
func GetCommands() []cli.Command {
	return []cli.Command{
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
		{
			Name:    "save",
			Aliases: []string{"sv"},
			Usage:   "Save a validated analysis",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "publish",
			Aliases: []string{"p"},
			Usage:   "Publish a saved analysis",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "supress",
			Aliases: []string{"su"},
			Usage:   "Supssspress a published analysis",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:    "search",
			Aliases: []string{"sr"},
			Usage:   "Suppress a published analysis",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
}
