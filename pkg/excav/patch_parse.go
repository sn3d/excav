package excav

import (
	"github.com/sn3d/excav/pkg/cast"
	"github.com/sn3d/excav/pkg/task"
	"github.com/sn3d/excav/pkg/task/append"
	"github.com/sn3d/excav/pkg/task/delete"
	"github.com/sn3d/excav/pkg/task/file"
	"github.com/sn3d/excav/pkg/task/put"
	"github.com/sn3d/excav/pkg/task/remove"
	"github.com/sn3d/excav/pkg/task/replace"
	"github.com/sn3d/excav/pkg/task/script"
	"gopkg.in/yaml.v2"
)

// parseYaml iterate over all records in YAML and fulfill the
// given Patch structure.
//
// we're decoding YAML as dynamic data. This part identify
// only name of the task and task's type. The rest of
// the YAML is forwarded data into appropriate parsing
// function.
//
// The YAML contains list of records like parameters, tasks etc.
// The type of the record is determined by first parameter.
// e.g. if record start with 'name' parameter, it's task.
func parseYaml(data []byte, out *Patch) error {
	var rootYml []interface{}
	var err error

	err = yaml.Unmarshal(data, &rootYml)
	if err != nil {
		return err
	}

	// forgive this ugly type casting, but it's because
	// dynamic nature of YAML and everything is interface{}
	for _, record := range rootYml {
		recData := cast.ToStrData(record)
		if hasParameter("name", recData) {
			parseTask(recData["name"].(string), recData, out)
			parseTaskMetadata(recData["name"].(string), recData, out)
		}
	}

	return nil
}

// The task record start with 'name' parameter and one of the further
// parameters is type of task e.g. 'append'. The YAML looks:
//
//	name: some-task
//	append:
//	   param1: xxxxx
//
// Every task might have metadata, some common parameters like 'only'
// etc. Those metadata are separated, because Task is interface
//
// Be careful. The type name shouldn't match to any metadata
// parameter!
func parseTask(name string, record map[string]interface{}, out *Patch) {
	for k, val := range record {
		var t task.Task
		var err error

		switch k {
		case "replace":
			t, err = replace.Parse(name, val)
		case "append":
			t, err = append.Parse(name, val)
		case "file":
			t, err = file.Parse(name, val)
		case "remove":
			t, err = remove.Parse(name, val)
		case "script":
			t, err = script.Parse(name, val)
		case "delete":
			t, err = delete.Parse(name, val)
		case "put":
			t, err = put.Parse(name, val)
		}

		if err != nil {
			break
		}

		if t != nil {
			out.Tasks[name] = t
		}
	}
}

// every task might have some metadata. Metadata are common parameters for
// every task like 'Only' etc.
func parseTaskMetadata(name string, record map[string]interface{}, out *Patch) {
	// parse metadata
	metadata := &task.Metadata{
		Only: cast.ToStr(record["only"]),
	}
	out.Metadata[name] = metadata
}

func hasParameter(param string, data map[string]interface{}) bool {
	for k, _ := range data {
		if k == param {
			return true
		}
	}
	return false
}
