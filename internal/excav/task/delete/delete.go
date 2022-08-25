package delete

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/sn3d/excav/api"
	"github.com/sn3d/excav/lib/cast"
)

type DeleteTask struct {
	Path        string
	BeginAnchor string
	EndAnchor   string
}

func Parse(name string, in interface{}) (api.Task, error) {
	task := DeleteTask{}
	for key, val := range cast.ToData(in) {
		switch cast.ToStr(key) {
		case "path":
			task.Path = cast.ToStr(val)
		case "begin":
			task.BeginAnchor = cast.ToStr(val)
		case "end":
			task.EndAnchor = cast.ToStr(val)
		}
	}
	return &task, nil
}

func (t *DeleteTask) Patch(dir string, params map[string]interface{}) error {
	// load file
	filePath := filepath.Join(dir, t.Path)
	fileInfo, _ := os.Stat(filePath)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	replaced := del(data, t.BeginAnchor, t.EndAnchor)

	// save new content & exit
	err = ioutil.WriteFile(filePath, replaced, fileInfo.Mode())
	return err
}

// This function iterate over given data and delete all content
// between begin and end anchors. Anchors are included in deletion.
//
// The function is repeating this operation and delete all
// content between all anchors, until reach the end of the data.
func del(data []byte, beginAnchor string, endAnchor string) []byte {

	var before []byte
	var result = make([]byte, 0)

	beginRe, err := regexp.Compile("(?m)" + beginAnchor)
	if err != nil {
		return nil
	}

	endRe, err := regexp.Compile("(?m)" + endAnchor)
	if err != nil {
		return nil
	}

	for {
		// everything before begin anchor
		beginIdx := beginRe.FindIndex(data)
		if beginIdx == nil {
			// this condition might be weird but I have to
			// return original data when no begin-anchor was
			// found. And this is it.
			result = append(result, data...)
			break
		}
		before = data[:beginIdx[0]]
		data = data[beginIdx[1]:]
		result = append(result, before...)

		// everything after end anchor
		endIdx := endRe.FindIndex(data)
		if endIdx == nil {
			break
		}
		data = data[endIdx[1]:]
	}

	return result
}
