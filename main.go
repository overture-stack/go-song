package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	app.Commands = GetCommands()

	app.Run(os.Args)
}
