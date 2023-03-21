package push

import (
	"fmt"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/excav/pkg/termui"
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:  "push",
	Usage: "Push all repositories in 'applied' state to remote and create MR",
	Flags: []cli.Flag{},

	// main entry point for 'apply'
	Action: func(ctx *cli.Context) error {
		cfg, err := excav.LoadConfiguration()
		if err != nil {
			return termui.CliError(ctx, fmt.Errorf("Cannot load the configuration. Check the ~/.config/excav/config.yaml"))
		}

		// initialize event dispatcher
		excav.RegisterListener(
			&termui.Listener{},
		)

		// ...and current bulk
		blk, err := excav.OpenBulk(cfg)
		if err != nil {
			return termui.CliError(ctx, fmt.Errorf("Cannot open the bulk"))
		}
		defer blk.Close()

		// pushing all 'applied' repos
		blk.PushAll()
		return nil
	},
}
