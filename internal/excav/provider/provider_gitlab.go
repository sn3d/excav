package provider

import (
	"fmt"
	"strings"

	"github.com/sn3d/excav/lib/gitlab"
)

type gitlabProvider struct {
	host  string
	token string
}

func (gl *gitlabProvider) GetRepositoryURL(repoName string) string {
	// get rid of first '/'
	if strings.HasPrefix(repoName, "/") {
		repoName = repoName[1:]
	}

	// get rid of last '/'
	host := gl.host
	if strings.HasSuffix(host, "/") {
		host = host[:len(host)-1]
	}

	return host + "/" + repoName
}

func (gl *gitlabProvider) GetToken() string {
	return gl.token
}

func (gl *gitlabProvider) GetUser() string {
	return ""
}

func (gl *gitlabProvider) CreateMergeRequest(repoName string, branch string, title string) (string, error) {

	// validation first. Better to do validation here then in GitLab.
	// It save you unnecessary request.
	if repoName == "" {
		return "", fmt.Errorf("repoName can't be blank")
	}

	if branch == "" {
		return "", fmt.Errorf("branch can't be blank")
	}

	if title == "" {
		return "", fmt.Errorf("title can't be blank")
	}

	client := gitlab.New(gl.host, gl.token)
	project := client.Projects().GetByPath(repoName)
	if project == nil {
		return "", fmt.Errorf("project %s not found", repoName)
	}

	mr := project.MergeRequests().Create(title, branch)
	if mr == nil {
		return "", fmt.Errorf("cannot create merge request")
	}

	return mr.Url, nil
}

func (gl *gitlabProvider) DeleteBranch(repoName string, branch string) error {
	client := gitlab.New(gl.host, gl.token)
	project := client.Projects().GetByPath(repoName)
	err := project.Branches().Delete(branch)
	if err != nil {
		return err
	}
	return nil
}
