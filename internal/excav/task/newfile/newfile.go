package newfile

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sn3d/excav/api"
	"github.com/sn3d/excav/lib/cast"
	"github.com/sn3d/excav/lib/template"
)

type NewFileTask struct {
	Src  string
	Dest string
	Mode os.FileMode
}

func Parse(name string, in interface{}) (api.Task, error) {
	task := NewFileTask{
		Mode: 0644,
	}

	for key, val := range cast.ToData(in) {
		switch cast.ToStr(key) {
		case "template":
			task.Src = cast.ToStr(val)
		case "path":
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

// This function apply 'add' task on repository directory
func (t *NewFileTask) Patch(dir string, params map[string]interface{}) error {

	dest := template.Subst(t.Dest, params)
	src := template.Subst(t.Src, params)

	// read data from source file
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	content := template.Subst(string(data), params)

	// ensure the directory for dest. file will exist
	destDir := filepath.Join(dir, filepath.Dir(dest))
	_, err = os.Stat(destDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(destDir, 0777)
		if err != nil {
			return err
		}
	}

	// write data to dest. file
	destFile := filepath.Join(dir, dest)
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
