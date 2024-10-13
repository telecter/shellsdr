package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func list(ctx *cli.Context) error {
	fmt.Println("Listing devices...")
	return nil
}

var List = &cli.Command{
	Name:   "list",
	Usage:  "List recognized SDR devices",
	Action: list,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "verbose",
			Usage:   "Provide additional information about devices",
			Aliases: []string{"v"},
		},
	},
}
