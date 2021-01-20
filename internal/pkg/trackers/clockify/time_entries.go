package clockify

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (c *Clockify) ListTimeEntries(from, to time.Time) ([]*trackers.TimeEntries, error) {
	if c.workspaceId == "" {
		return nil, errors.New("this method requires a workspace id")
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
		return nil, errors.New("invalid status code")
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
