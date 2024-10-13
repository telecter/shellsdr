package main

import (
	"log"
	"os"

	"github.com/telecter/shellsdr/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:  "shellsdr",
		Usage: "Control your RTL-SDR in a command line app.",
		Commands: []*cli.Command{
			cmd.List,
		},
		Flags: []cli.Flag{},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
