package excav

import (
	"github.com/sn3d/excav/api"
	"github.com/sn3d/excav/internal/excav/task"
	"github.com/sn3d/excav/lib/cast"
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
//     name: some-task
//     append:
//        param1: xxxxx
//
// Every task might have metadata, some common parameters like 'only'
// etc. Those metadata are separated, because Task is interface
//
// Be careful. The type name shouldn't match to any metadata
// parameter!
func parseTask(name string, record map[string]interface{}, out *Patch) {
	for k, val := range record {
		parser := task.Parsers[k]
		if parser != nil {
			task, _ := parser(name, val)
			if task != nil {
				out.Tasks[name] = task
				break
			}
		}
	}
}

// every task might have some metadata. Metadata are common parameters for
// every task like 'Only' etc.
func parseTaskMetadata(name string, record map[string]interface{}, out *Patch) {
	// parse metadata
	metadata := &api.TaskMetadata{
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
