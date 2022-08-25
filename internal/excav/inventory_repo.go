package excav

import (
	"strings"
)

type Repository struct {
	// Path of repository is in GitLab `/fodler/folder/repo` or in GitLab `/org/repo`
	Name string `yaml:"repository"`

	// not every repository have 'main' as default branch. Some legacy repos
	// might have 'master' as default branch
	DefaultBranch string `yaml:"default_branch"`

	// tags help you identify and select list of
	// repositories
	Tags []string `yaml:"tags,flow"`

	// repository might have some parameters they are
	// passed into templates
	RepoParams Params `yaml:"params,omitempty"`
}

// get the trimmed name ready to
// use. We're not using direct access,
// because Name might have weird characters
func (r *Repository) GetName() string {
	if strings.HasSuffix(r.Name, "/") {
		return r.Name[:len(r.Name)-1]
	}
	return r.Name
}

func GetRepositoryNames(repos []*Repository) []string {
	repoNames := make([]string, len(repos))
	for i, repo := range repos {
		repoNames[i] = repo.Name
	}
	return repoNames
}

func (repo *Repository) HasTags(tags ...string) bool {
	for _, tag := range tags {
		if !repo.HasTag(tag) {
			return false
		}
	}
	return true
}

func (repo *Repository) HasTag(tag string) bool {
	for _, t := range repo.Tags {
		if t == tag {
			return true
		}
	}
	return false
}
