package clockify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (c *Clockify) ListTimeEntries(from, to time.Time) ([]*trackers.TimeEntries, error) {
	if c.workspaceId == "" || c.userId == "" {
		return nil, errors.New("this method requires a workspace id and a user id")
	}
	u := c.getUrl(fmt.Sprintf("/workspaces/%s/user/%s/time-entries", c.workspaceId, c.userId))
	q := u.Query()
	q.Add("start", from.Format(time.RFC3339))
	q.Add("end", to.Format(time.RFC3339))
	u.RawQuery = q.Encode()
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
	var entries []timeEntries
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, err
	}
	genericEntries := make([]*trackers.TimeEntries, 0, len(entries))
	for _, entry := range entries {
		if entry.WorkspaceId == c.workspaceId {
			genericEntries = append(genericEntries, entry.ToGeneric())
		}
	}
	return genericEntries, nil
}

func (c *Clockify) CreateTimeEntry(timeEntry *trackers.TimeEntries) (*trackers.TimeEntries, error) {
	if c.workspaceId == "" {
		return nil, errors.New("this method requires a workspace id")
	}
	u := c.getUrl(fmt.Sprintf("/workspaces/%s/time-entries", c.workspaceId))
	createEntity := &createTimeEntries{
		Start:       timeEntry.Start,
		End:         timeEntry.End,
		Description: timeEntry.Description,
		ProjectId:   timeEntry.ProjectId,
	}
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(&createEntity); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), &body)
	if err != nil {
		return nil, err
	}
	c.addAuthentication(req)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}
	var createdEntity timeEntries
	if err := json.NewDecoder(resp.Body).Decode(&createdEntity); err != nil {
		return nil, err
	}
	return createdEntity.ToGeneric(), nil
}
