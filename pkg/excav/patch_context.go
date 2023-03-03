package excav

import (
	"errors"
	"path/filepath"

	"github.com/sn3d/excav/pkg/git"
	"github.com/sn3d/excav/pkg/log"
	"github.com/sn3d/excav/pkg/provider"
)

type PhaseType string

const (
	PhaseNew     PhaseType = ""
	PhaseApplied PhaseType = "applied"
	PhasePushed  PhaseType = "pushed"
)

// PatchContext hold the repository and metadata related to current patching.
// When Bulk represent the group of repositories you apply patch, the
// PatchContext represent one repository you're currently patching
type PatchContext struct {
	RepoName string `yaml:"repo_name"`

	// from which branch is new branch for patch created, by default it's 'main'
	DefaultBranch string `yaml:"default_branch"`

	// this is a created branch (from default) where is patch applied
	Branch string `yaml:"branch"`

	Phase PhaseType `yaml:"phase"`

	// Description is provided by user fist time, when bulk is created
	// (see CreateNew)
	Description string `yaml:"descr"`

	// This field is populated after changes are pushed (see bulk Push())
	MergeRequestURL string `yaml:"merge_request"`

	// This field indicates error state. You can check the state
	// via IsError func.
	ErrorMsg string `yaml:"error_msg"`

	// This field holds all parameters and values for patch
	AllParams Params `yaml:"params"`

	bulkDir Directory `yaml:"-"`

	// ref. to provider
	prvd provider.Provider
}

func (ctx *PatchContext) IsError() bool {
	return ctx.ErrorMsg != ""
}

// apply the given p to repository and commit the changes.
func (ctx *PatchContext) Apply(p *Patch, message string) error {
	ctx.Phase = PhaseApplied

	// open right branch
	repo, err := ctx.cloneOrOpen()
	if err != nil {
		ctx.handleError(err)
		return err
	}

	err = repo.Checkout(ctx.Branch)
	if err != nil {
		ctx.handleError(err)
		return err
	}

	err = p.Apply(ctx.getRepositoryDir().String(), ctx.AllParams.ToMap())
	if err != nil {
		ctx.handleError(err)
		return err
	}

	// commit changes
	err = repo.CommitAll(message)
	if err != nil {
		ctx.handleError(err)
		return err

	}

	return nil
}

func (ctx *PatchContext) Push() error {
	ctx.Phase = PhasePushed
	repo, err := ctx.cloneOrOpen()
	if err != nil {
		return err
	}

	log.Debug("pushing branch", "branch", ctx.Branch, "user", ctx.prvd.GetUser())
	err = repo.PushNewBranch(ctx.prvd.GetUser(), ctx.prvd.GetToken(), ctx.Branch)
	if err != nil {
		log.Error("cannot push new branch", err)
		ctx.handleError(errors.New("Cannot push new branch. Probably branch already exist"))
		return err
	}

	log.Debug("create MR/PR")
	ctx.MergeRequestURL, err = ctx.prvd.CreateMergeRequest(ctx.RepoName, ctx.Branch, ctx.Description)
	if err != nil {
		log.Error("cannot create merge request for branch", err)
		ctx.handleError(err)
		return err
	}

	ctx.ErrorMsg = ""

	return nil
}

func (ctx *PatchContext) Diff() string {
	repo, err := ctx.open()
	if err != nil {
		return ""
	}
	return repo.Diff(ctx.DefaultBranch, ctx.Branch)
}

func (ctx *PatchContext) handleError(err error) {
	ctx.ErrorMsg = err.Error()
}

func (ctx *PatchContext) cloneOrOpen() (*git.Repository, error) {
	var repo *git.Repository
	var err error

	// check if repo is already cloned or not
	repoDir := ctx.getRepositoryDir()
	if repoDir.IsNotExist() {
		// clone repository
		repoURL := ctx.prvd.GetRepositoryURL(ctx.RepoName)
		log.Debug("cloning repository", "repoUrl", repoURL, "dir", string(repoDir), "branch", ctx.DefaultBranch, "user", ctx.prvd.GetUser())
		repo, err = git.Clone(repoURL, string(repoDir), ctx.DefaultBranch, ctx.prvd.GetUser(), ctx.prvd.GetToken())
		return repo, err
	} else {
		log.Debug("open repository", "dir", string(repoDir))
		return ctx.open()
	}
}

func (ctx *PatchContext) open() (*git.Repository, error) {
	repoDir := ctx.getRepositoryDir()
	return git.Open(string(repoDir))
}

func (ctx *PatchContext) getRepositoryDir() Directory {
	repoDir := ctx.bulkDir.Subdir(ctx.RepoName)
	absRepoDir, err := filepath.Abs(string(repoDir))
	if err != nil {
		return ""
	}
	return Directory(absRepoDir)
}
