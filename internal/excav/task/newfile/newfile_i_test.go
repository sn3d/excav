package newfile_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/sn3d/excav/internal/excav"
	"github.com/sn3d/excav/lib/testdata"
)

func Test_AddTask(t *testing.T) {
	testdata.Prepare()

	// given patch with 'add' task
	patchDir := testdata.AbsPath("patch-newfile")
	p, _ := excav.OpenPatch(patchDir)

	// when we apply patch
	params := map[string]interface{}{
		"name":    "steve",
		"message": "Hello!",
	}

	repoDir := testdata.AbsPath("repo")
	err := p.Apply(repoDir, params)
	if err != nil {
		t.FailNow()
	}

	// then the file must be added to 'path/to/' folder
	_, err = os.Stat(filepath.Join(repoDir, "path/to/readme-steve.txt"))
	if os.IsNotExist(err) {
		t.FailNow()
	}
}
