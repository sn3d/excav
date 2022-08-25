package gitlab

import (
	"encoding/json"
	"strings"
)

type Projects struct {
	client *Client
}

func (projects *Projects) Get(id int) *Project {
	data, status := projects.client.get("/api/v4/projects/%d", id)
	if status != 200 {
		return nil
	}

	proj := &Project{
		client: projects.client,
	}

	err := json.Unmarshal([]byte(data), proj)
	if err != nil {
		return nil
	}

	return proj
}

func (projects *Projects) GetByPath(path string) *Project {

	// remove first '/'
	if path[0] == '/' {
		path = path[1:]
	}

	pathEncoded := strings.ReplaceAll(path, "/", "%2F")
	data, status := projects.client.get("/api/v4/projects/%s", pathEncoded)
	if status != 200 {
		return nil
	}

	proj := &Project{
		client: projects.client,
	}

	err := json.Unmarshal([]byte(data), proj)
	if err != nil {
		return nil
	}

	return proj
}
