package task

import (
	"github.com/sn3d/excav/api"
	"github.com/sn3d/excav/internal/excav/task/append"
	"github.com/sn3d/excav/internal/excav/task/delete"
	"github.com/sn3d/excav/internal/excav/task/file"
	"github.com/sn3d/excav/internal/excav/task/put"
	"github.com/sn3d/excav/internal/excav/task/remove"
	"github.com/sn3d/excav/internal/excav/task/replace"
	"github.com/sn3d/excav/internal/excav/task/script"
)

type ParseTaskFunc func(name string, in interface{}) (api.Task, error)

// Here are registered all parsers for all supported tasks.
// The key is string used in YAML and value is parsing function
var Parsers = map[string]ParseTaskFunc{
	"replace": replace.Parse,
	"append":  append.Parse,
	"file":    file.Parse,
	"remove":  remove.Parse,
	"script":  script.Parse,
	"delete":  delete.Parse,
	"put":     put.Parse,
}
