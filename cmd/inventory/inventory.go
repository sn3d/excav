package inventory

import (
	"fmt"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/excav/pkg/termui"
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:  "inventory",
	Usage: "Manipulate with inventory or repositories",
	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "tag",
			Aliases: []string{"t"},
		},
		&cli.StringFlag{
			Name:    "inventory",
			Aliases: []string{"i"},
		},
	},

	Action: func(ctx *cli.Context) error {
		inventoryFile := ctx.String("inventory")
		if inventoryFile == "" {
			inventoryFile = "./inventory.yaml"
		}

		inv, err := excav.OpenInventory(inventoryFile)
		if err != nil {
			fmt.Printf("error: %v", err)
			return err
		}

		repos := inv.GetByTags(ctx.StringSlice("tag")...)
		for _, repo := range repos {
			termui.PrintRepository(repo.Name, repo.Tags, repo.RepoParams)
		}

		return nil
	},
}
