package clockify

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (c *Clockify) HasWorkspace() bool {
	return true
}

func (c *Clockify) ListWorkspaces() ([]*trackers.Workspace, error) {
	u := c.getUrl("/workspaces")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	c.addAuthentication(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	var workspaces []*workspace
	if err := json.NewDecoder(resp.Body).Decode(&workspaces); err != nil {
		return nil, err
	}
	genericWorkspaces := make([]*trackers.Workspace, 0, len(workspaces))
	for _, workspace := range workspaces {
		genericWorkspaces = append(genericWorkspaces, workspace.ToGeneric())
	}
	return genericWorkspaces, nil
}
