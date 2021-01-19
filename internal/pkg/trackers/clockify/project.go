package clockify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (c *Clockify) ListProjects() ([]*trackers.Project, error) {
	if c.workspaceId == "" {
		return nil, errors.New("this method requires a workspace id")
	}
	u := c.getUrl(fmt.Sprintf("/workspaces/%s/projects", c.workspaceId))
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
