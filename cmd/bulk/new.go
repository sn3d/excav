package bulk

import (
	"errors"
	"fmt"

	"github.com/sn3d/excav/internal/excav"
	"github.com/sn3d/excav/internal/termui"
	"github.com/urfave/cli/v2"
)

var newCmd = &cli.Command{
	Name:  "new",
	Usage: "Create new bulk and set it as current. The old one isn't discarded.",

	// main entry point for 'state ls'
	Action: func(ctx *cli.Context) error {

		cfg, err := excav.LoadConfiguration()
		if err != nil {
			return termui.CliError(ctx, errors.New("Cannot load configuration. Check the ~/.config/excav/config.yaml"))
		}

		bulk := excav.NewBulk(cfg)
		if err != nil {
			return cli.Exit("Cannot open current bulk", 1)
		}

		fmt.Println(bulk.ID)
		return nil
	},
}
