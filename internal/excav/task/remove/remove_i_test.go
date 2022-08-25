package remove_test

import (
	"os"
	"testing"

	"github.com/sn3d/excav/internal/excav"
	"github.com/sn3d/excav/lib/testdata"
)

func Test_RemoveTask(t *testing.T) {

	// given patch with 'remove' task
	testdata.Prepare()
	patchDir := testdata.AbsPath("patch-remove")

	p, err := excav.OpenPatch(patchDir)
	if err != nil {
		t.FailNow()
	}

	// when we apply patch
	repoDir := testdata.AbsPath("repo")
	p.Apply(repoDir, nil)

	// then the files must be deleted
	_, err = os.Stat(testdata.AbsPath("repo/subdir/file1.txt"))
	if os.IsExist(err) {
		t.Error("the file1.txt exist! It should be deleted")
		t.FailNow()
	}

	_, err = os.Stat(testdata.AbsPath("repo/subdir/file2.txt"))
	if os.IsExist(err) {
		t.Error("the file2.txt exists! It should be deleted")
	}
}
