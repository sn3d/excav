package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type githubProvider struct {
	gitHost string
	apiHost string
	token   string
	user    string
}

func (gh *githubProvider) GetRepositoryURL(repoName string) string {
	return gh.gitHost + "/" + repoName
}

func (gh *githubProvider) GetToken() string {
	return gh.token
}
func (gh *githubProvider) GetUser() string {
	return gh.user
}

func (gh *githubProvider) CreateMergeRequest(repoName string, branch string, commitMsg string) (string, error) {
	// Let's cut the '/' prefix in repo name
	// repository might start with '/' (e.g. /sn3d/repo)
	// but GitHub API have problems with //sn3d/repo URLs.
	if strings.HasPrefix(repoName, "/") {
		repoName = repoName[1:]
	}

	defaultBranch := gh.getDefaultBranch(repoName)
	url := fmt.Sprintf("%s/repos/%s/pulls", gh.apiHost, repoName)
	payload := fmt.Sprintf("{ \"title\": \"%s\", \"head\":\"%s\", \"base\": \"%s\" }", commitMsg, branch, defaultBranch)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+gh.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != 201 {
		return "", fmt.Errorf("status code is %d and we expect 200", resp.StatusCode)
	}

	var respBody struct {
		HtmlUrl string `json:"html_url"`
	}

	err = json.Unmarshal(buf.Bytes(), &respBody)
	if err != nil {
		return "", fmt.Errorf("cannot parse json response %v", err)
	}

	return respBody.HtmlUrl, nil
}

func (gh *githubProvider) DeleteBranch(repoName string, branch string) error {
	return errors.New("unimplemented function")
}

func (gh *githubProvider) getDefaultBranch(repoName string) string {
	url := gh.apiHost + "/repos/" + repoName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ""
	}
	req.Header.Add("Content-Type", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", "token "+gh.token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != 200 {
		return ""
	}

	var respBody struct {
		DefaultBranch string `json:"default_branch"`
	}
	err = json.Unmarshal(buf.Bytes(), &respBody)
	if err != nil {
		return ""
	}

	return respBody.DefaultBranch
}
