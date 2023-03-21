package initialize

import (
	"os"
	"path/filepath"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/excav/pkg/termui"
	"github.com/urfave/cli/v2"
)

var Cmd = &cli.Command{
	Name:  "init",
	Usage: "Create config directories and generate global configuration",

	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "gitlab",
		},
		&cli.StringFlag{
			Name: "gitlab-token",
		},
		&cli.StringFlag{
			Name: "github",
		},
		&cli.StringFlag{
			Name: "github-api",
		},
		&cli.StringFlag{
			Name: "github-token",
		},
	},

	// main entry point for 'apply'
	Action: func(ctx *cli.Context) error {
		cfg := excav.Configuration{
			GitLabHost:    ctx.String("gitlab"),
			GitLabToken:   ctx.String("gitlab-token"),
			GitHubHost:    ctx.String("github"),
			GitHubApiHost: ctx.String("github-api"),
			GitHubToken:   ctx.String("github-token"),
		}

		provider := termui.Select("What provider do you want to use?", "gitlab", "github")
		switch provider {
		case "gitlab":
			askGitLab(&cfg)
		case "github":
			askGitHub(&cfg)
		}
		//askGitLab(&cfg)

		// save configuration to config file
		userHome, _ := os.UserHomeDir()

		configDir := filepath.Join(userHome, ".config", "excav")
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			os.MkdirAll(configDir, 0777)
		}

		configFile := filepath.Join(configDir, "config.yaml")
		err := cfg.SaveToYaml(configFile)
		if err != nil {
			return termui.CliError(ctx, err)

		}

		termui.Info("Configuration saved to %s", configFile)
		return nil
	},
}

func askGitLab(cfg *excav.Configuration) {
	if cfg.GitLabHost == "" {
		cfg.GitLabHost = termui.Ask("What's the GitLab host (default: https://gitlab.com)?")
		if cfg.GitLabHost == "" {
			cfg.GitLabHost = "https://gitlab.com"
		}
	}

	if cfg.GitLabToken == "" {
		cfg.GitLabToken = termui.Ask("What's your GitLab personal access token?")
	}
}

func askGitHub(cfg *excav.Configuration) {
	if cfg.GitHubHost == "" {
		cfg.GitHubHost = termui.Ask("What's the GitHub host (default: https://github.com)?")
		if cfg.GitHubHost == "" {
			cfg.GitHubHost = "https://github.com"
		}
	}

	if cfg.GitHubApiHost == "" {
		cfg.GitHubApiHost = termui.Ask("What's the GitHub API host (default: https://api.github.com)?")
		if cfg.GitHubApiHost == "" {
			cfg.GitHubApiHost = "https://api.github.com"
		}
	}

	if cfg.GitHubUser == "" {
		cfg.GitHubUser = termui.Ask("What's your GitHub user?")
	}

	if cfg.GitHubToken == "" {
		cfg.GitHubToken = termui.Ask("What's your GitHub personal access token?")
	}
}
