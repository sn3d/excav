package put

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sn3d/excav/api"
	"github.com/sn3d/excav/lib/cast"
	"github.com/sn3d/excav/lib/log"
	"github.com/sn3d/excav/lib/template"
)

type PutTask struct {
	Path string

	// regexp of anchor. Anchor is place in the file
	// where the content will be placed. (e.g. '// +excav:anchor')
	Anchor string

	// string that will be appended. It's good if you don't want
	// to provide template file, you need only simple string
	Content string

	// path to template file what content will be used for
	// appending.
	Template string

	// expression ensure the task will be executed
	// only once.
	//
	// Let's imagine we want to add 'server-123' into file.
	// If we will create put task for this content, and we will
	// run it twice, we will have twice 'server-123'.
	//
	// By setting WhenNot to 'server-123', we will ensure, the task
	// will be applied only WHEN file NOT contains 'server-123' at all.
	WhenNot string
}

func Parse(name string, in interface{}) (api.Task, error) {
	task := PutTask{}
	for key, val := range cast.ToData(in) {
		switch cast.ToStr(key) {
		case "path":
			task.Path = cast.ToStr(val)
		case "anchor":
			task.Anchor = cast.ToStr(val)
		case "content":
			task.Content = cast.ToStr(val)
		case "template":
			task.Template = cast.ToStr(val)
		case "when_not":
			task.WhenNot = cast.ToStr(val)
		}
	}
	return &task, nil
}

func (t *PutTask) Patch(dir string, params map[string]interface{}) error {
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

func (t *PutTask) PatchFile(file string, params map[string]interface{}) error {
	var err error
	var sb strings.Builder

	log.Debug("[put] patching file", "file", file)
	content, _ := t.getContent()

	// open the file we're going to patch
	fileInfo, _ := os.Stat(file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// skip it if it's already applied
	if isAlreadyApplied(t.WhenNot, data) {
		return nil
	}

	// get all anchors
	expr := "(?m)^(\\s+)" + t.Anchor + "$"
	re, err := regexp.Compile(expr)
	if err != nil {
		return err
	}

	allIndexes := re.FindAllSubmatchIndex(data, -1)
	pos := 0
	for _, idx := range allIndexes {
		// we found anchor
		log.Debug("[put] found anchor")

		// write the original text before anchor
		sb.Write(data[pos:idx[0]])

		// write anchor
		sb.Write(data[idx[0]:idx[1]])

		// write intended content/template after anchor
		paramsWithGroups := addGroupsToParams(params, idx, data)
		finalContent := template.Subst(content, paramsWithGroups)
		indentedContent := addIndent(finalContent, data[idx[2]:idx[3]])
		sb.WriteString(indentedContent)

		pos = idx[1]
	}

	// write rest of the original text
	sb.Write(data[pos:])

	// save new content into file
	err = ioutil.WriteFile(file, []byte(sb.String()), fileInfo.Mode())
	return err
}

// returns you real content that will be placed. All
// template placeholders are replaced by values from
// params etc.
func (t *PutTask) getContent() (string, error) {
	if len(t.Content) > 0 {
		return t.Content, nil
	} else {
		data, err := ioutil.ReadFile(t.Template)
		if err != nil {
			return "", err
		}
		return string(data), err
	}
}

func addIndent(content string, indent []byte) string {
	var sb strings.Builder

	splitted := strings.Split(content, "\n")
	for _, line := range splitted {
		// and end-of-line if it's needed (if it's missing)
		if len(line) == 0 {
			sb.WriteRune('\n')
		} else {
			if line[len(line)-1] != 10 {
				sb.WriteRune('\n')
			}
		}

		if len(indent) > 0 && indent[0] == 10 {
			indent = indent[1:]
		}

		sb.Write(indent)
		sb.WriteString(line)
	}

	return sb.String()
}

func addGroupsToParams(params map[string]interface{}, groups []int, data []byte) map[string]interface{} {
	result := map[string]interface{}{}

	// copy original params first
	for k, v := range params {
		result[k] = v
	}

	// add regex groups to param 'groups'
	groupsParam := []string{}

	// let's strip the first 2 records because first is value
	// itself and second is an special indent group added by
	// PutTask.
	//
	// People know groups are indexed by 1, not 0 because first
	// group is regexp itself. Here, we're beginning with 0.
	// I hope that will not cause confusion.
	for i := 4; i < len(groups); i = i + 2 {
		val := string(data[groups[i]:groups[i+1]])
		groupsParam = append(groupsParam, val)
	}
	result["groups"] = groupsParam

	return result
}

func isAlreadyApplied(condition string, data []byte) bool {
	if condition == "" {
		return false
	}

	expr := "(?m)^(\\s+)" + condition + "$"
	re, err := regexp.Compile(expr)
	if err != nil {
		log.Error("cannot compile 'when_not' expression", err)
		return true
	}

	m := re.Match(data)
	return m
}
