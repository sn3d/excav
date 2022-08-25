package main

import (
	"os"

	"github.com/sn3d/excav/cmd"
	"github.com/sn3d/excav/internal/termui"
	"github.com/sn3d/excav/lib/log"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:     "excav",
		Before:   before,
		Commands: cmd.Commands,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "debug",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		termui.ShowError(err)
	}
}

func before(ctx *cli.Context) error {

	// when you specify "--debug", the logger is enabled,
	// otherwise is disabled and only standard messages are
	// print
	if ctx.Bool("debug") {
		log.EnableDebug()
	}

	return nil
}
