package gitlab

import (
	"encoding/json"
)

type Pipelines struct {
	client    *Client
	projectID int
}

func (ps *Pipelines) Get(id int) *Pipeline {
	data, code := ps.client.get("/api/v4/projects/%d/pipelines/%d", ps.projectID, id)
	if code != 200 {
		return nil
	}

	var pipeline Pipeline
	err := json.Unmarshal([]byte(data), &pipeline)
	if err != nil {
		return nil
	}

	pipeline.client = ps.client
	pipeline.projectID = ps.projectID

	return &pipeline
}

// GetAll returns all pipelines they're matching to filter. e.g.
// If filter is 'status=running'. This filter is appended into URL. For that
// it must follow format 'key1=val1&key2=val2'. For more info about
// supported filter parameters see https://docs.gitlab.com/ee/api/pipelines.html#list-project-pipelines
func (ps *Pipelines) GetAll(filter string) []*Pipeline {
	data, code := ps.client.get("/api/v4/projects/%d/pipelines?%s", ps.projectID, filter)
	if code != 200 {
		return []*Pipeline{}
	}

	var pipelines []*Pipeline
	err := json.Unmarshal([]byte(data), &pipelines)
	if err != nil {
		return []*Pipeline{}
	}

	// enrich the client and project ID
	for _, pipeln := range pipelines {
		pipeln.projectID = ps.projectID
		pipeln.client = ps.client
	}

	return pipelines
}
