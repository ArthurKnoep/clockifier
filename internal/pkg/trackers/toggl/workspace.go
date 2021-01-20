package toggl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (t *Toggl) HasWorkspace() bool {
	return true
}

func (t *Toggl) ListWorkspaces() ([]*trackers.Workspace, error) {
	u := t.getUrl("/workspaces")
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	t.addAuthentication(req)
	resp, err := t.httpClient.Do(req)
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
