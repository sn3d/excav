package excav

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sn3d/excav/pkg/log"
	"github.com/sn3d/excav/pkg/task"
)

// Patch contains one or more tasks they're applied to
// some directory.
type Patch struct {
	currentDir string
	Params     Params
	Tasks      map[string]task.Task
	Metadata   map[string]*task.Metadata
}

// Open and returns patch. It's read and parse 'patch.yaml' for
// given directory.
func OpenPatch(dir string) (*Patch, error) {

	patch := &Patch{
		currentDir: dir,
		Params:     make(Params),
		Tasks:      make(map[string]task.Task),
		Metadata:   make(map[string]*task.Metadata),
	}

	mainFile := filepath.Join(dir, "patch.yaml")
	data, err := ioutil.ReadFile(mainFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s/patch.yaml", dir)
	}

	err = parseYaml(data, patch)
	if err != nil {
		return nil, fmt.Errorf("cannot parse YAML %s/patch.yaml: %v", dir, err)
	}

	return patch, nil
}

// Apply patch to directory(absolute path), where is repository cloned.
// If some task failed, the apply terminate process with error of task.
//
// This function also emmit events like TaskStarted and TaskEnd
// into global dispatcher.
func (p *Patch) Apply(repoDir string, params map[string]interface{}) error {
	log.Debug("apply patch")

	originDir, _ := os.Getwd()

	defer os.Chdir(originDir)
	err := os.Chdir(p.currentDir)
	if err != nil {
		return err
	}

	for taskName, task := range p.Tasks {
		log.Debug("apply task", "task", taskName)
		dispatcher.Notify(TaskStarted{Task: taskName})

		err := task.Patch(repoDir, params)
		if err != nil {
			dispatcher.Notify(TaskEnd{Task: taskName, Error: err})
			log.Error("cannot apply task", err)
			return err
		}

		dispatcher.Notify(TaskEnd{Task: taskName})
	}

	log.Debug("patch applied")
	return nil
}
