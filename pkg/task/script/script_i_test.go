package script_test

import (
	"io/ioutil"
	"testing"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/testdata"
)

func Test_ScriptTask(t *testing.T) {
	testdata.Setup()

	// given patch with replace tasks
	p, err := excav.OpenPatch(testdata.Abs("patch-script"))
	if err != nil {
		t.FailNow()
	}

	// when we apply patch
	err = p.Apply(testdata.Abs("repo"), nil)
	if err != nil {
		t.FailNow()
	}

	// then python script create the script.txt
	text, _ := ioutil.ReadFile(testdata.Abs("repo/script.txt"))
	if string(text) != "Hello World" {
		t.FailNow()
	}
}
