package toggl

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (t *Toggl) ListTimeEntries(from, to time.Time) ([]*trackers.TimeEntries, error) {
	if t.workspaceId == "" {
		return nil, errors.New("this method requires a workspace id")
	}
	u := t.getUrl("/me/time_entries")
	q := u.Query()
	q.Add("start_date", from.Format(time.RFC3339))
	q.Add("end_date", to.Format(time.RFC3339))
	u.RawQuery = q.Encode()
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
	var entries []timeEntries
	if err := json.NewDecoder(resp.Body).Decode(&entries); err != nil {
		return nil, err
	}
	genericEntries := make([]*trackers.TimeEntries, 0, len(entries))
	for _, entry := range entries {
		if strconv.Itoa(entry.WorkspaceId) == t.workspaceId && entry.Duration >= 0 {
			genericEntries = append(genericEntries, entry.ToGeneric())
		}
	}
	return genericEntries, nil
}

func (t *Toggl) CreateTimeEntry(timeEntry *trackers.TimeEntries) (*trackers.TimeEntries, error) {
	panic("not implemented")
}
