package put_test

import (
	"fmt"
	"testing"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/excav/pkg/testdata"
)

func Test_PutTask(t *testing.T) {
	testdata.Setup()

	fmt.Println(testdata.String())

	p, err := excav.OpenPatch(testdata.AbsPath("patch"))
	if err != nil {
		t.FailNow()
	}

	absPath := testdata.AbsPath("repo")
	err = p.Apply(absPath, nil)
	if err != nil {
		t.FailNow()
	}

	// apply again
	err = p.Apply(absPath, nil)
	if err != nil {
		t.FailNow()
	}

	// then 'file1.g' have to match 'file1.expected':w
	isMatching := testdata.CompareFiles("repo/file.1", "repo/file-1.expected")
	if !isMatching {
		t.FailNow()
	}
}
