package toggl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (t *Toggl) ListProjects() ([]*trackers.Project, error) {
	if t.workspaceId == "" {
		return nil, errors.New("this method requires a workspace id")
	}
	u := t.getUrl(fmt.Sprintf("/workspaces/%s/projects", t.workspaceId))
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
		return nil, errors.New("invalid status code")
	}
	var projects []project
	if err := json.NewDecoder(resp.Body).Decode(&projects); err != nil {
		return nil, err
	}
	genericProjects := make([]*trackers.Project, 0, len(projects))
	for _, project := range projects {
		genericProjects = append(genericProjects, project.ToGeneric())
	}
	return genericProjects, nil
}

