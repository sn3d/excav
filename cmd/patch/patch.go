package patch

import (
	"fmt"
	"os"

	"github.com/sn3d/excav/internal/excav"
	"github.com/sn3d/excav/internal/termui"
	"github.com/urfave/cli/v2"
)

type Options struct {
	Patch  string
	Dir    string
	Params excav.Params
}

var Cmd = &cli.Command{
	Name:      "patch",
	Usage:     "Patching of concrete directory",
	ArgsUsage: "[patch] [dir]",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "param",
			Aliases: []string{"p"},
		},
	},

	// main entry point for 'patch'
	Action: func(ctx *cli.Context) error {
		opts := Options{}
		var err error

		opts.Patch = ctx.Args().Get(0)
		if opts.Patch == "" {
			return termui.CliError(ctx, fmt.Errorf("patch is mandatory"))
		}

		opts.Dir = ctx.Args().Get(1)
		if opts.Dir == "" {
			opts.Dir, err = os.Getwd()
			if err != nil {
				return termui.CliError(ctx, fmt.Errorf("cannot use current working dir"))
			}
		}

		opts.Params = excav.StringSliceToParams(ctx.StringSlice("param"))

		err = applyPatch(&opts)
		if err != nil {
			return termui.CliError(ctx, err)
		}

		return nil
	},
}

func applyPatch(opts *Options) error {
	var err error
	var pth *excav.Patch

	pth, err = excav.OpenPatch(opts.Patch)
	if err != nil {
		return err
	}

	err = pth.Apply(opts.Dir, opts.Params.ToMap())
	if err != nil {
		return err
	}

	return nil
}
