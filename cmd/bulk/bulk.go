package bulk

import "github.com/urfave/cli/v2"

var Cmd = &cli.Command{
	Name:  "bulk",
	Usage: "Show and manipulate with bulk",
	Subcommands: []*cli.Command{
		newCmd,
	},
}
