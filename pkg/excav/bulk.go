package excav

import (
	"fmt"

	"github.com/sn3d/excav/pkg/api"
	"github.com/sn3d/excav/pkg/log"
	"github.com/sn3d/excav/pkg/provider"
)

type NewBulkOpts struct {
	Repos        []*Repository
	Branch       string
	Description  string
	GlobalParams Params
}

type Bulk struct {
	ID string

	dir      Directory
	contexts []*PatchContext
	cfg      *Configuration
	prvd     provider.Provider
}

//------------------------------------------------------------------------------
// Public API
//------------------------------------------------------------------------------

// funcion create new empty bulk, folder of this bulk and
// set it as current.
func NewBulk(cfg *Configuration) *Bulk {

	bulk := Bulk{
		ID:       randomName(),
		contexts: make([]*PatchContext, 0),
		cfg:      cfg,
		prvd:     newProvider(cfg),
	}

	// create new bulk dir. with random name
	bulk.dir = Directory(cfg.WorkspaceDir).Subdir(bulk.ID)
	if bulk.dir.IsNotExist() {
		bulk.dir.Mkdir()
	}

	cfg.SetCurrentBulk(bulk.ID)

	return &bulk
}

func OpenBulk(cfg *Configuration) (*Bulk, error) {

	// initialize bulk
	bulk := Bulk{
		ID:   cfg.GetCurrentBulk(),
		dir:  cfg.GetCurrentBulkDir(),
		cfg:  cfg,
		prvd: newProvider(cfg),
	}

	log.Debug("open bulk dir", "dir", bulk.dir)
	if bulk.dir.IsNotExist() {
		return nil, fmt.Errorf("Empty or non-existing bulk")
	}

	// load all contexts in buldDir
	contexts := make([]*PatchContext, 0)
	bulk.dir.FindBySuffix(".yaml", func(file string) {
		ctx, err := loadContext(file)
		if err == nil {
			ctx.bulkDir = bulk.dir
			ctx.prvd = bulk.prvd
			contexts = append(contexts, ctx)
		} else {
			log.Error("cannot load context", err)
		}
	})
	bulk.contexts = contexts

	return &bulk, nil
}

func (b *Bulk) AddRepositories(repos []*Repository, globalParams Params, commitMsg string, branch string) {
	for _, repo := range repos {
		b.AddRepository(repo, globalParams, commitMsg, branch)
	}
}

func (b *Bulk) AddRepository(repo *Repository, globalParams Params, mergeRequestDescr string, branch string) {

	// check if repository is already added
	for _, ctx := range b.contexts {
		if ctx.RepoName == repo.GetName() {
			return // skip it
		}
	}

	ctx := &PatchContext{
		RepoName:      repo.GetName(),
		DefaultBranch: repo.DefaultBranch,
		Branch:        branch,
		Description:   mergeRequestDescr,
		AllParams:     MergeParams(repo.RepoParams, globalParams),
		bulkDir:       b.dir,
		prvd:          b.prvd,
	}

	b.contexts = append(b.contexts, ctx)
}

// Apply patch to only given repository in bulk
func (b *Bulk) Apply(repoName string, p *Patch, message string) {
	for _, ctx := range b.contexts {
		if ctx.RepoName == repoName {
			dispatcher.Notify(api.PatchingStarted{
				Repo: ctx.RepoName,
			})

			ctx.Apply(p, message)

			dispatcher.Notify(api.PatchApplied{
				Repo:      ctx.RepoName,
				Branch:    ctx.Branch,
				CommitMsg: message,
				ErrorMsg:  ctx.ErrorMsg,
			})
		}
	}
}

func (b *Bulk) PushAll() {
	for _, ctx := range b.contexts {
		ctx.Push()
		dispatcher.Notify(api.Pushed{
			Repo:            ctx.RepoName,
			MergeRequestURL: ctx.MergeRequestURL,
			ErrorMsg:        ctx.ErrorMsg,
		})
	}
}

// Close save the state
func (b *Bulk) Close() error {
	for _, context := range b.contexts {
		if err := saveContext(b.dir, context); err != nil {
			//TODO: handle error
			dispatcher.Notify(api.RepoError{
				Repo:     context.RepoName,
				ErrorMsg: "error saving state of repo",
				Error:    err,
			})
		}
	}
	return nil
}

// Len returns you number of patch contexts
func (b *Bulk) Len() int {
	return len(b.contexts)
}

// ForEach iterate over contexts and call f function
// for every context. If f returns some error (not null),
// the iteration immediately stop and forward this error
// to you
func (b *Bulk) ForEach(f func(ctx *PatchContext) error) error {
	for _, ctx := range b.contexts {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Bulk) Get(repoName string) *PatchContext {
	for _, ctx := range b.contexts {
		if ctx.RepoName == repoName {
			return ctx
		}
	}
	return nil
}

// Discarding the Bulk. That means:
//   - delete bulk dir content
//   - delette branches if they're created
func (b *Bulk) Discard() error {
	b.dir.Remove()

	// discarding merge requests
	for _, ctx := range b.contexts {
		if ctx.MergeRequestURL != "" { // discard only pushed MRs
			err := b.prvd.DeleteBranch(ctx.RepoName, ctx.Branch)
			dispatcher.Notify(api.RepoDiscarded{
				Repo:  ctx.RepoName,
				Error: err,
			})
		}
	}

	// unset current bulk
	b.cfg.SetCurrentBulk("")

	dispatcher.Notify(api.BulkDiscarded{})
	return nil
}

func newProvider(cfg *Configuration) provider.Provider {
	prvd := provider.New(provider.Options{
		Provider: provider.Typ(cfg.GetProviderType()),
		Host:     cfg.GetProviderHost(),
		ApiHost:  cfg.GetProviderApiHost(),
		Token:    cfg.GetProviderToken(),
		User:     cfg.GetProviderUser(),
	})
	return prvd
}
