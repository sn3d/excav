package diff

import (
	"fmt"

	"github.com/sn3d/excav/internal/excav"
	"github.com/sn3d/excav/internal/termui"
	"github.com/urfave/cli/v2"
)

type Options struct {
	InventoryFile string
}

var Cmd = &cli.Command{
	Name:      "diff",
	Usage:     "Show the diff after patching",
	ArgsUsage: "<repository>",

	Flags: []cli.Flag{},

	// main entry point for 'diff'
	Action: func(ctx *cli.Context) error {
		repoName := ctx.Args().First()

		cfg, err := excav.LoadConfiguration()
		if err != nil {
			return termui.CliError(ctx, fmt.Errorf("Cannot load configuration. Check the ~/.config/excav/config.yaml"))
		}

		// open bulk
		blk, err := excav.OpenBulk(cfg)
		if err != nil {
			return termui.CliError(ctx, fmt.Errorf("Cannot open bulk. Check the .excav folder"))
		}

		// diff all repos or concrete repo - it depends if user provide
		// repo name as argument
		if repoName == "" {
			blk.ForEach(printDiff)
		} else {
			ctx := blk.Get(repoName)
			if ctx != nil {
				printDiff(ctx)
			}
		}
		return nil
	},
}

func printDiff(ctx *excav.PatchContext) error {
	diff := ctx.Diff()
	fmt.Printf("\n%s: %s\n", termui.Magenta("Repository"), termui.BrightWhite(ctx.RepoName))
	fmt.Println(diff)
	return nil
}
