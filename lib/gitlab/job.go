package gitlab

type Job struct {
	// private
	client    *Client
	projectID int

	// JSON attributes
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func (j *Job) Trace() string {
	data, status := j.client.get("/api/v4/projects/%d/jobs/%d/trace", j.projectID, j.ID)
	if status != 200 {
		return ""
	}
	return data
}
