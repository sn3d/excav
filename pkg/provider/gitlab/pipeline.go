package gitlab

import "encoding/json"

type Pipeline struct {
	client    *Client
	projectID int
	ID        int    `json:"id"`
	Status    string `json:"status"`
}

func (p *Pipeline) Jobs() []*Job {
	data, status := p.client.get("/api/v4/projects/%d/pipelines/%d/jobs", p.projectID, p.ID)
	if status != 200 {
		return []*Job{}
	}

	var jobs []*Job
	err := json.Unmarshal([]byte(data), &jobs)
	if err != nil {
		return []*Job{}
	}

	// enrich client
	for _, job := range jobs {
		job.client = p.client
		job.projectID = p.projectID
	}

	return jobs
}
