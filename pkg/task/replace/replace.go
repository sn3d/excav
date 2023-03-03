package replace

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sn3d/excav/api"
	"github.com/sn3d/excav/pkg/cast"
	"github.com/sn3d/excav/pkg/log"
	"github.com/sn3d/excav/pkg/template"
)

type ReplaceTask struct {
	Path     string
	Regexp   string
	Text     string
	Replace  string
	Template string
}

func Parse(name string, in interface{}) (api.Task, error) {
	task := ReplaceTask{}
	for key, val := range cast.ToData(in) {
		switch cast.ToStr(key) {
		case "path":
			task.Path = cast.ToStr(val)
		case "text":
			task.Text = cast.ToStr(val)
		case "regexp":
			task.Regexp = cast.ToStr(val)
		case "replace":
			task.Replace = cast.ToStr(val)
		case "template":
			task.Template = cast.ToStr(val)
		}
	}
	return &task, nil
}

func (t *ReplaceTask) Patch(dir string, params map[string]interface{}) error {
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if info != nil {
			if !info.IsDir() {
				relative := p[len(dir)+1:]
				pathExpr := template.Subst(t.Path, params)
				isMatching, _ := path.Match(pathExpr, relative)
				if isMatching {
					err := t.PatchFile(p, params)
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	})
	return nil
}

func (t *ReplaceTask) PatchFile(path string, params map[string]interface{}) error {
	log.Debug("[replace] patching file", "file", path)

	// open the file we want to patch
	fileInfo, _ := os.Stat(path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// get content we will use as replacement
	content, err := t.getContent(params)
	if err != nil {
		return err
	}

	// do replacing
	var replaced []byte
	if len(t.Regexp) > 0 {
		// assuming we replace by regexp.
		finalRegexp := template.Subst(t.Regexp, params)
		log.Debug("[replace] replace by regexp", "regexp", finalRegexp)

		re, err := regexp.Compile(finalRegexp)
		if err != nil {
			return err
		}
		replaced = re.ReplaceAll(data, []byte(content))
	} else {
		// assuming we replace text
		replaceText := template.Subst(t.Text, params)
		log.Debug("[replace] replace by text", "text", replaceText)

		newText := strings.ReplaceAll(string(data), replaceText, content)
		replaced = []byte(newText)
	}

	// save new content of file
	err = ioutil.WriteFile(path, replaced, fileInfo.Mode())
	if err != nil {
		return err
	}

	return nil
}

func (t *ReplaceTask) getContent(params map[string]interface{}) (string, error) {
	var content string = t.Replace
	if t.Template != "" {
		data, err := ioutil.ReadFile(t.Template)
		if err != nil {
			return "", err
		}
		content = string(data)
	}

	content = template.Subst(content, params)
	return content, nil
}
