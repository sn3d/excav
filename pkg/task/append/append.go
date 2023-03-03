package append

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sn3d/excav/pkg/api"
	"github.com/sn3d/excav/pkg/cast"
	"github.com/sn3d/excav/pkg/log"
	"github.com/sn3d/excav/pkg/template"
)

type ModeType string

const (
	ModeAppendBegin  ModeType = "append-begin"
	ModeAppendEnd    ModeType = "append-end"
	ModeAppendBefore ModeType = "append-before"
	ModeAppendAfter  ModeType = "append-after"
)

type AppendTask struct {
	// exact relative path to file which will be
	// patched.
	Path string

	Mode   ModeType
	Anchor string

	// string that will be appended. It's good if you don't want
	// to provide template file, you need only simple string
	Content string

	// path to template file what content will be used for
	// appending.
	Template string

	// it's used as conditional reg.exp. The patch is applied only when
	// file's content matching to this reg exp. If condition starts with 'not ',
	// then 'not ' is cutout and patch is applied when file's content not
	// matching to regexp.
	When string
}

func Parse(name string, in interface{}) (api.Task, error) {
	task := AppendTask{}
	for key, val := range cast.ToData(in) {
		switch cast.ToStr(key) {
		case "path":
			task.Path = cast.ToStr(val)
		case "mode":
			task.Mode = ModeType(cast.ToStr(val))
		case "anchor":
			task.Anchor = cast.ToStr(val)
		case "content":
			task.Content = cast.ToStr(val)
		case "template":
			task.Template = cast.ToStr(val)
		case "when":
			task.When = cast.ToStr(val)
		}
	}
	return &task, nil
}

func (t *AppendTask) Patch(dir string, params map[string]interface{}) error {

	// get the file we want to patch
	path := template.Subst(t.Path, params)
	dir, _ = filepath.Abs(dir)
	file := filepath.Join(dir, path)

	err := t.PatchFile(file, params)
	if err != nil {
		log.Error("cannot append to file:", err)
		return err
	}

	return nil
}

func (t *AppendTask) PatchFile(file string, params map[string]interface{}) error {
	log.Debug("[append] patching file", "file", file)

	// get the content we want to apply
	content, err := t.getContent()
	if err != nil {
		return err
	}
	content = template.Subst(content, params)

	fileInfo, _ := os.Stat(file)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	// check the 'When' condition (it it's set)
	var pass = true
	if t.When != "" {
		if strings.HasPrefix(t.When, "not ") {
			pass = !when(t.When[len("not "):], data)
		} else {
			pass = when(t.When, data)
		}
	}
	if !pass {
		return nil
	}

	// do insertion by mode
	var replaced []byte
	switch t.Mode {
	case ModeAppendBegin:
		replaced = appendAtBegin(content, data)
	case ModeAppendEnd:
		replaced = appendAtEnd(content, data)
	case ModeAppendBefore:
		replaced = appendBefore(t.Anchor, content, data)
	case ModeAppendAfter:
		replaced = appendAfter(t.Anchor, content, data)
	}

	// save new content
	err = ioutil.WriteFile(file, replaced, fileInfo.Mode())
	if err != nil {
		return err
	}
	return nil
}

func (t *AppendTask) getContent() (string, error) {
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

func when(condition string, data []byte) bool {
	if condition == "" {
		return true
	}
	re, err := regexp.Compile("(?m)" + condition)
	if err != nil {
		return false
	}
	m := re.Match(data)
	return m
}

func appendAtBegin(content string, data []byte) []byte {
	return append([]byte(content), data...)
}

func appendAtEnd(content string, data []byte) []byte {
	return append(data, []byte(content)...)
}

func appendBefore(anchor string, content string, data []byte) []byte {
	re, err := regexp.Compile("(?m)^(\\s+)" + anchor + "$")
	if err != nil {
		return []byte{}
	}

	// get the position of first regexp match - first anchor
	// append content before anchor
	idx := re.FindSubmatchIndex(data)
	if idx == nil {
		return data
	}

	var sb strings.Builder

	// append beginning
	sb.Write(data[:idx[0]])

	// put spaces
	sb.Write(data[idx[2]:idx[3]])

	// put content before anchor
	sb.WriteString(content)
	sb.WriteRune('\n')

	// append rest
	sb.Write(data[idx[0]:])

	return []byte(sb.String())
}

func appendAfter(anchor string, content string, data []byte) []byte {
	re, err := regexp.Compile("(?m)^(\\s+)" + anchor + "$")
	if err != nil {
		return []byte{}
	}

	// get the position of first regexp match - this is anchor
	// append content after anchor
	idx := re.FindSubmatchIndex(data)
	if idx == nil {
		return data
	}

	var sb strings.Builder

	// append beginning
	sb.Write(data[:idx[1]])

	// put new line and spaces
	sb.WriteRune('\n')
	sb.Write(data[idx[2]:idx[3]])

	// put content
	sb.WriteString(content)

	// append rest
	sb.Write(data[idx[1]:])

	return []byte(sb.String())
}
