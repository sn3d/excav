package gitlab

type Project struct {
	client *Client
	ID     int      `json:"id"`
	Path   string   `json:"path_with_namespace"`
	Topics []string `json:"tag_list"`
}

func (project *Project) MergeRequests() *MergeRequests {
	return &MergeRequests{
		client:    project.client,
		projectID: project.ID,
	}
}

func (project *Project) Branches() *Branches {
	return &Branches{
		client:    project.client,
		projectID: project.ID,
	}
}

func (project *Project) Pipelines() *Pipelines {
	return &Pipelines{
		client:    project.client,
		projectID: project.ID,
	}
}
