package show

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sn3d/excav/internal/excav"
	"github.com/sn3d/excav/internal/termui"
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:  "show",
	Usage: "Show state of the bulk",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "output",
			Aliases:     []string{"o"},
			Usage:       "output format: text/json. Default is text",
			DefaultText: "text",
		},
	},

	// main entry point for 'state ls'
	Action: func(ctx *cli.Context) error {
		cfg, err := excav.LoadConfiguration()
		if err != nil {
			return termui.CliError(ctx, errors.New("Cannot load configuration. Check the ~/.config/excav/config.yaml"))
		}

		blk, err := excav.OpenBulk(cfg)
		if err != nil {
			return cli.Exit("Cannot open current bulk", 1)
		}

		if ctx.String("output") == "json" {
			describeAsJson(blk)
		} else {
			describeAsText(blk)
		}

		return nil
	},
}

func describeAsText(blk *excav.Bulk) error {

	blk.ForEach(func(ctx *excav.PatchContext) error {

		fmt.Printf("%s: %s\n", termui.Magenta("Repository"), termui.BrightWhite(ctx.RepoName))
		switch ctx.Phase {
		case excav.PhaseApplied:
			if ctx.IsError() {
				fmt.Printf("   [%s] Applied\n", termui.Red(termui.XMark))
				fmt.Printf("       %s: %v\n", termui.Red("Error"), ctx.ErrorMsg)
				fmt.Printf("   [ ] Pushed\n")
			} else {
				fmt.Printf("   [%s] Applied\n", termui.Green(termui.CheckMark))
				fmt.Printf("   [ ] Pushed\n")
			}
		case excav.PhasePushed:
			if ctx.IsError() {
				fmt.Printf("   [%s] Applied\n", termui.Green(termui.CheckMark))
				fmt.Printf("   [%s] Pushed\n", termui.Red(termui.XMark))
				fmt.Printf("       %s: %v\n", termui.Red("Error"), ctx.ErrorMsg)
			} else {
				fmt.Printf("   [%s] Applied\n", termui.Green(termui.CheckMark))
				fmt.Printf("   [%s] Pushed\n", termui.Green(termui.CheckMark))
				fmt.Printf("       MR: %s\n", termui.Hyperlink(ctx.MergeRequestURL, ctx.MergeRequestURL))
			}
		}
		return nil
	})

	return nil
}

func describeAsJson(blk *excav.Bulk) {
	var builder strings.Builder
	builder.WriteRune('[')

	isFirst := true
	blk.ForEach(func(ctx *excav.PatchContext) error {
		// ensure the ','
		if isFirst {
			isFirst = false
		} else {
			builder.WriteRune(',')
		}

		// start of JSON object
		builder.WriteRune('{')

		// 'phase'
		builder.WriteString(" \"phase\":")
		builder.WriteString(" \"")
		builder.WriteString(string(ctx.Phase))
		builder.WriteString("\",")

		if ctx.Phase == excav.PhasePushed {
			builder.WriteString(" \"merge_request\":")
			builder.WriteString(" \"")
			builder.WriteString(ctx.MergeRequestURL)
			builder.WriteString("\",")
		}

		// 'repository'
		builder.WriteString(" \"repository\":")
		builder.WriteString(" \"")
		builder.WriteString(ctx.RepoName)
		builder.WriteString("\"")

		// end JSON object
		builder.WriteRune('}')
		return nil
	})

	builder.WriteRune(']')
	fmt.Print(builder.String())
}
