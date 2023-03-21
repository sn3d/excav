package script

import (
	"os/exec"
	"strings"
	"text/template"

	"github.com/sn3d/excav/pkg/cast"
	"github.com/sn3d/excav/pkg/task"
)

type ScriptTask struct {
	Commands []string
}

type scriptData struct {
	RepositoryDir string
}

func Parse(name string, in interface{}) (task.Task, error) {
	data := in.([]interface{})
	task := ScriptTask{
		Commands: cast.ToStrArr(data),
	}
	return &task, nil
}

func (t *ScriptTask) Patch(repoDir string, params map[string]interface{}) error {
	data := &scriptData{
		RepositoryDir: repoDir,
	}

	for _, cmd := range t.Commands {
		err := runCommand(cmd, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func runCommand(cmd string, data *scriptData) error {
	// replace placeholders
	var cmdStr strings.Builder
	tmpl, err := template.New("cmd").Parse(cmd)
	if err != nil {
		return err
	}

	err = tmpl.Execute(&cmdStr, data)
	if err != nil {
		return err
	}

	// execute it
	cmdWithArgs := strings.Fields(cmdStr.String())
	err = exec.Command(cmdWithArgs[0], cmdWithArgs[1:]...).Run()
	return err
}
