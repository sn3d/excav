package delete_test

import (
	"testing"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/excav/pkg/testdata"
)

// scenario:
//   - given repository with some code
//   - when we apply patch with delete tas
//   - then parts in code.txt should be deleted
func Test_DeleteTask(t *testing.T) {
	testdata.Setup()

	repoDir := testdata.AbsPath("repo")
	patchDir := testdata.AbsPath("patch-delete")

	p, err := excav.OpenPatch(patchDir)
	if err != nil {
		t.FailNow()
	}

	err = p.Apply(repoDir, map[string]interface{}{})
	if err != nil {
		t.FailNow()
	}

	equal := testdata.CompareFiles("repo/code.txt", "expected.txt")
	if !equal {
		t.FailNow()
	}
}
