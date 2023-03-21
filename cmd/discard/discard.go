package discard

import (
	"errors"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/excav/pkg/termui"
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:  "discard",
	Usage: "Discard current bulk (workspace, existing merge requests).",

	// main entry point for 'state ls'
	Action: func(ctx *cli.Context) error {
		cfg, err := excav.LoadConfiguration()
		if err != nil {
			return termui.CliError(ctx, errors.New("Cannot load configuration. Check the ~/.config/excav/config.yaml"))
		}

		blk, err := excav.OpenBulk(cfg)
		if err != nil {
			return termui.CliError(ctx, err)
		}

		err = blk.Discard()
		if err != nil {
			return termui.CliError(ctx, err)
		}

		return nil
	},
}
