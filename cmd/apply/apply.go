package apply

import (
	"errors"
	"fmt"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/excav/pkg/termui"
	"github.com/urfave/cli/v2"
)

type Options struct {
	Patch         string
	Repositories  []string
	Tags          []string
	InventoryFile string
	CommitMsg     string
	Branch        string
	Params        []string
}

var Cmd = &cli.Command{
	Name:      "apply",
	Usage:     "Apply patch to repository or repositories",
	ArgsUsage: "[patch]",

	Flags: []cli.Flag{
		&cli.StringSliceFlag{
			Name:    "tag",
			Aliases: []string{"t"},
		},
		&cli.StringSliceFlag{
			Name:    "repository",
			Aliases: []string{"r"},
		},
		&cli.StringFlag{
			Name:    "inventory",
			Aliases: []string{"i"},
		},
		&cli.StringFlag{
			Name:    "commit",
			Aliases: []string{"m"},
		},
		&cli.StringFlag{
			Name:    "branch",
			Aliases: []string{"b"},
		},
		&cli.StringSliceFlag{
			Name:    "param",
			Aliases: []string{"p"},
		},
		&cli.BoolFlag{
			Name:    "reuse",
			Aliases: []string{"u"},
			Usage:   "reuse current bulk and not create a new one",
		},
		&cli.BoolFlag{
			Name:    "no-interactive",
			Aliases: []string{"n"},
			Usage:   "non interactive mode disable all interactive prompts",
		},
	},

	// main entry point for 'apply'
	Action: func(ctx *cli.Context) error {
		err := applyAction(ctx)
		if err != nil {
			return termui.CliError(ctx, err)
		}

		return nil
	},
}

// Real 'apply' function. When all options are parsed
// correctly, config is loaded, we can proceed to action itself.
func applyAction(ctx *cli.Context) error {
	var (
		patchDir       = ctx.Args().First()
		tags           = ctx.StringSlice("tag")
		repositories   = ctx.StringSlice("repository")
		inventoryFile  = ctx.String("inventory")
		commitMsg      = ctx.String("commit")
		branch         = ctx.String("branch")
		globalParams   = excav.StringSliceToParams(ctx.StringSlice("param"))
		reuseBulk      = ctx.Bool("reuse")
		nonInteractive = ctx.Bool("no-interactive")
	)

	// validation & load config
	if patchDir == "" {
		return errors.New("Patch is mandatory")
	}

	config, err := excav.LoadConfiguration()
	if err != nil {
		return errors.New("Cannot load the configuration. Check the ~/.config/excav/config.yaml")
	}

	// we need load patch first - what to apply
	p, err := excav.OpenPatch(patchDir)
	if err != nil {
		return fmt.Errorf("Cannot load patch from %s", patchDir)
	}

	// excav emit events. You can react on these events
	// via registered listener implementation.
	// In this case, it's terminal UI
	excav.RegisterListener(&termui.Listener{})

	// select repositories we want to patch
	inv, _ := excav.OpenInventory(inventoryFile)
	selectedRepos := make([]*excav.Repository, 0)

	if len(repositories) > 0 {
		for _, repoName := range repositories {
			repo := inv.Get(repoName)
			if repo != nil {
				selectedRepos = append(selectedRepos, repo)
			}
		}
	}

	if len(tags) > 0 {
		taggedRepos := inv.GetByTags(tags...)
		selectedRepos = append(selectedRepos, taggedRepos...)
	}

	if len(selectedRepos) == 0 {
		selectedRepos = inv.GetAll()
	}

	// ask if user want to patch selected repositories
	repoNames := excav.GetRepositoryNames(selectedRepos)
	termui.PrintRepositories(repoNames)

	if !nonInteractive {
		if !termui.ConfirmApply() {
			return nil
		}
	}

	if commitMsg == "" {
		commitMsg = termui.Ask("Enter a commit message")
	}

	if branch == "" {
		branch = termui.Ask("Enter branch name")
	}

	// create a new bulk or reuse old one
	var blk *excav.Bulk
	if reuseBulk {
		blk, err = excav.OpenBulk(config)
		if err != nil {
			return fmt.Errorf("cannot open current bulk: %v", err)
		}
	} else {
		blk = excav.NewBulk(config)
	}

	// populate repositories
	blk.AddRepositories(selectedRepos, globalParams, commitMsg, branch)

	// during apply, the events like status, errors,
	// they will be dispatch to disp
	defer blk.Close()

	// let's apply to selected repos
	for _, selectedRepo := range selectedRepos {
		blk.Apply(selectedRepo.GetName(), p, commitMsg)
	}

	// and print some status
	fmt.Println(termui.Green("\nThe patch was applied"))
	fmt.Println("You can check the result with " + termui.BrightWhite("excav diff"))
	fmt.Println("Or continue pushing changes with " + termui.BrightWhite("excav push\n"))
	return nil
}
