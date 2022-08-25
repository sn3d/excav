package api

// Task is main abstraction. Tasks like 'replace' etc. are implementing
// this abstraction.
type Task interface {
	Patch(dir string, params map[string]interface{}) error
}

type TaskMetadata struct {
	Only string
}
