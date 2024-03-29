package replace_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sn3d/excav/pkg/excav"
	"github.com/sn3d/testdata"
)

// scenario: replace tag with content from template file
//
//	when: we apply patch 'patch-replace-tmplt' to 'example-replace' repository
//	then: 'TODO' text need to be changed by content from 'template.txt'
func Test_ReplaceTask(t *testing.T) {
	var err error
	var p *excav.Patch

	testdata.Setup()
	repoDir := testdata.Abs("repo")

	p, err = excav.OpenPatch(testdata.Abs("patch-replace-tmplt"))
	if err != nil {
		t.Errorf("Error read patch %s", err)
	}

	err = p.Apply(repoDir, nil)
	if err != nil {
		t.Errorf("Error in apply of patch %s", err)
	}

	// then...
	text, _ := ioutil.ReadFile(testdata.Abs("repo/file3.txt"))
	if strings.Contains(string(text), "TODO: text here") {
		t.Errorf("the +tag wasn't replaced. Check if replacing works correctly and +tag is in file3.txt")
	}

	if !strings.Contains(string(text), "Hello Template!") {
		t.Errorf("the text from template.txt wasn't replaced. Check if file was loaded correctly")
	}

}
