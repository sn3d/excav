package excav

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sn3d/excav/pkg/dir"
	"gopkg.in/yaml.v2"
)

func loadContext(path string) (*PatchContext, error) {
	ctx := &PatchContext{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ctx, err
	}

	err = yaml.Unmarshal(data, &ctx)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func saveContext(bulkDir dir.Directory, ctx *PatchContext) error {
	if ctx.RepoName == "" {
		return errors.New("cannot save context, repo name is empty string")
	}

	data, err := yaml.Marshal(ctx)
	if err != nil {
		return err
	}

	path := composeContextFileName(bulkDir, ctx.RepoName)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// this function returns you absolute path to file where
// state if PatchContext is stored
func composeContextFileName(workspace dir.Directory, repoName string) string {

	// normalize name '/org/my-repo' to 'org-my-repo'
	normalized := repoName
	if normalized[0] == '/' {
		normalized = normalized[1:]
	}
	normalized = strings.ReplaceAll(normalized, "/", "-")

	// returns abs. path to patch metadata file
	return workspace.File(normalized + ".yaml")
}
