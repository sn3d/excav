package main

import (
	"os"

	"github.com/sn3d/excav/cmd/apply"
	"github.com/sn3d/excav/cmd/bulk"
	"github.com/sn3d/excav/cmd/diff"
	"github.com/sn3d/excav/cmd/discard"
	"github.com/sn3d/excav/cmd/initialize"
	"github.com/sn3d/excav/cmd/inventory"
	"github.com/sn3d/excav/cmd/patch"
	"github.com/sn3d/excav/cmd/push"
	"github.com/sn3d/excav/cmd/show"

	"github.com/sn3d/excav/internal/termui"
	"github.com/sn3d/excav/lib/log"
	"github.com/urfave/cli/v2"
)

// version is set by goreleaser, via -ldflags="-X 'main.version=...'".
var version = "development"

func main() {

	app := &cli.App{
		Name:    "excav",
		Before:  before,
		Version: version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "debug",
			},
		},
		Commands: []*cli.Command{
			initialize.Cmd,
			apply.Cmd,
			bulk.Cmd,
			push.Cmd,
			diff.Cmd,
			show.Cmd,
			discard.Cmd,
			patch.Cmd,
			inventory.Cmd,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		termui.ShowError(err)
	}
}

// this function is executed before any subcommand and
// handle the --debug option
func before(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		log.EnableDebug()
	}
	return nil
}
