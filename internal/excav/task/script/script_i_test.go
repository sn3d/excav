package script_test

import (
	"io/ioutil"
	"testing"

	"github.com/sn3d/excav/internal/excav"
	"github.com/sn3d/excav/lib/testdata"
)

func Test_ScriptTask(t *testing.T) {
	testdata.Prepare()

	// given patch with replace tasks
	p, err := excav.OpenPatch(testdata.AbsPath("patch-script"))
	if err != nil {
		t.FailNow()
	}

	// when we apply patch
	err = p.Apply(testdata.AbsPath("repo"), nil)
	if err != nil {
		t.FailNow()
	}

	// then python script create the script.txt
	text, _ := ioutil.ReadFile(testdata.AbsPath("repo/script.txt"))
	if string(text) != "Hello World" {
		t.FailNow()
	}
}
