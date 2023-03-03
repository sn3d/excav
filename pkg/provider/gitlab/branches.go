package gitlab

import (
	"fmt"
	"strings"
)

type Branches struct {
	client    *Client
	projectID int
}

func (b *Branches) Delete(branch string) error {
	// we need to escape the '/' in branch name
	escBranch := strings.ReplaceAll(branch, "/", "%2f")
	_, code := b.client.del("/api/v4/projects/%d/repository/branches/%s", b.projectID, escBranch)
	if code != 204 {
		return fmt.Errorf("Error in delete branch %s", branch)
	}
	return nil
}
