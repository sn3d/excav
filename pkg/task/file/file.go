package file

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sn3d/excav/pkg/cast"
	"github.com/sn3d/excav/pkg/task"
	"github.com/sn3d/excav/pkg/template"
)

type FileTask struct {
	Src  string
	Dest string
	Mode os.FileMode
}

// This function is called when YAML parser found 'file' task.
func Parse(name string, in interface{}) (task.Task, error) {
	task := FileTask{
		Mode: 0644,
	}

	for key, val := range cast.ToData(in) {
		switch cast.ToStr(key) {
		case "src":
			task.Src = cast.ToStr(val)
		case "dest":
			task.Dest = cast.ToStr(val)
		case "mode":
			mode := cast.ToUint(val)
			if mode != 0 {
				task.Mode = os.FileMode(mode)
			}
		}
	}
	return &task, nil
}

// This function apply task on repository. The dir argument is full path to repo
// directory and params is map of all parameters passed to task.
func (t *FileTask) Patch(dir string, params map[string]interface{}) error {

	destFile := filepath.Join(dir, template.Subst(t.Dest, params))
	src := template.Subst(t.Src, params)

	// read data from source file
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	content := template.Subst(string(data), params)

	// ensure the directory for dest. file will exist
	destDir := filepath.Dir(destFile)
	_, err = os.Stat(destDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(destDir, 0777)
		if err != nil {
			return err
		}
	}

	// write data to dest. file
	f, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, t.Mode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte(content))
	if err != nil {
		return err
	}

	return nil
}
