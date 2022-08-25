package gitlab

import (
	"encoding/json"
	"fmt"
)

type MergeRequests struct {
	client    *Client
	projectID int
}

type MergeRequest struct {
	client    *Client
	projectID int
	ID        int    `json:"id"`
	IID       int    `json:"iid"`
	Url       string `json:"web_url"`
}

//------------------------------------------------------------------------------
// Merge requests collection functions
//------------------------------------------------------------------------------

func (mrs *MergeRequests) Create(title string, branch string) *MergeRequest {
	dataIn := fmt.Sprintf(`
		{ "title": "%s", 
          "source_branch": "%s", 
          "target_branch": "master", 
          "remove_source_branch": "true" 
        }`, title, branch)

	dataOut, status := mrs.client.post(dataIn, "/api/v4/projects/%d/merge_requests", mrs.projectID)
	if status != 201 {
		return nil
	}

	mr := &MergeRequest{
		client:    mrs.client,
		projectID: mrs.projectID,
	}

	err := json.Unmarshal([]byte(dataOut), mr)
	if err != nil {
		return nil
	}

	return mr
}
