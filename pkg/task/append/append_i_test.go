package append_test

import (
	"testing"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/testdata"
)

func Test_AppendTask(t *testing.T) {
	testdata.Setup()

	repoDir := testdata.Abs("repo")
	patchDir := testdata.Abs("patch")

	p, err := excav.OpenPatch(patchDir)
	if err != nil {
		t.FailNow()
	}

	err = p.Apply(repoDir, map[string]interface{}{})
	if err != nil {
		t.FailNow()
	}

	// then 'file.1' have to match 'file-1.expected'
	isMatching := testdata.CompareFiles("repo/file.1", "repo/file-1.expected")
	if !isMatching {
		t.FailNow()
	}

}
