package git

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"go.uber.org/zap"
)

// Via CloneOptions you can set various parameters for cloning like
// type of auth or fast cloning. If you wish standard cloning, you can pass
// empty CloneOptions into Clone
type CloneOptions struct {
	// Fast clone is cloning only single 'master' branch
	// with depth set to 1. It's useful when you don't need whole
	// history
	//FastClone bool

	// we're clonning only one branch - master on main. It depends on
	// provider (GitHub/GitLab)
	Branch string

	// personal access token to provider like GitHub or GitLab
	Token string
}

type Repository struct {
	Dir string

	//auth transport.AuthMethod
	repo *git.Repository
}

//------------------------------------------------------------------------------
// Public API
//------------------------------------------------------------------------------

// Open existing repository
func Open(dir string) (*Repository, error) {
	// open the directory
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return nil, err
	}

	out := &Repository{
		Dir:  dir,
		repo: repo,
	}

	return out, nil
}

// Clone do cloning with basic auth. This is useful for services like
// GitHub or GitLab where we use token as password.
//
// If branch is empty, the 'main' branch will be used
//
// The Clone function do fast clone. That means it clones only one branch
// and depth 1 commit
func Clone(src, dest, branch, authUser string, authToken string) (*Repository, error) {
	if branch == "" {
		branch = "main"
	}

	gitOpts := &git.CloneOptions{
		URL:           src,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		Auth: &http.BasicAuth{
			Username: authUser,
			Password: authToken,
		},

		// Fast clone
		Depth:        1,
		SingleBranch: true,
	}

	// cloning
	zap.L().Sugar().Debugw("Clone", "url", gitOpts.URL, "to", dest)
	repo, err := git.PlainClone(dest, false, gitOpts)
	if err != nil {
		return nil, err
	}

	// create repository instance
	out := &Repository{
		Dir:  dest,
		repo: repo,
	}

	return out, err
}

// Checkout the branch. If local branch doesn't exist,
// it's created new one.
func (r *Repository) Checkout(branch string) error {
	var err error

	if branch == "" {
		return errors.New("cannot checkout empty branch")
	}

	wtree, _ := r.repo.Worktree()
	branchRef := plumbing.ReferenceName("refs/heads/" + branch)

	err = wtree.Checkout(&git.CheckoutOptions{
		Branch: branchRef,
		Create: false,
		Force:  false,
	})

	// create it if failed
	if err != nil {
		err = wtree.Checkout(&git.CheckoutOptions{
			Branch: branchRef,
			Create: true,
			Force:  false,
		})
	}
	return err
}

// CommitAll is analogy to "git add --all && commit -m <msg>"
func (r *Repository) CommitAll(msg string) error {
	wtree, err := r.repo.Worktree()
	if err != nil {
		return err
	}

	// add all files to index
	err = wtree.AddWithOptions(&git.AddOptions{All: true})
	if err != nil {
		return err
	}

	wtree.Status()

	// commit
	_, err = wtree.Commit(msg, &git.CommitOptions{})
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) PushNewBranch(authUser string, authToken string, branch string) error {
	// we want to push new branch, this is bit tricky because
	// we need to create also connection between local and remote
	upstreamRef := plumbing.ReferenceName("refs/heads/" + branch)
	err := r.repo.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: authUser,
			Password: authToken,
		},
		RemoteName: "origin",
		RefSpecs: []config.RefSpec{
			config.RefSpec(upstreamRef + ":" + upstreamRef),
		},
	})
	return err
}

func (r *Repository) Diff(branch1, branch2 string) string {
	b1Hash, _ := r.repo.ResolveRevision(plumbing.Revision("refs/heads/" + branch1))
	b1Commit, _ := r.repo.CommitObject(*b1Hash)

	b2Hash, _ := r.repo.ResolveRevision(plumbing.Revision("refs/heads/" + branch2))
	b2Commit, _ := r.repo.CommitObject(*b2Hash)

	patch, _ := b1Commit.Patch(b2Commit)

	return patch.String()
}
