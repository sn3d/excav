package remove

import (
	"os"
	"path/filepath"

	"github.com/sn3d/excav/pkg/cast"
	"github.com/sn3d/excav/pkg/task"
)

type RemoveTask struct {
	Files []string
}

func Parse(name string, in interface{}) (task.Task, error) {
	task := RemoveTask{}
	for key, val := range cast.ToData(in) {
		switch cast.ToStr(key) {
		case "files":
			task.Files = cast.ToStrArr(val)
		}
	}
	return &task, nil
}

// This function apply 'remove' task on repository directory
func (t *RemoveTask) Patch(dir string, params map[string]interface{}) error {
	for _, file := range t.Files {
		err := os.RemoveAll(filepath.Join(dir, file))
		if err != nil {
			return err
		}
	}
	return nil
}
