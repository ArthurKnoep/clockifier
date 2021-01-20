package clockify

import (
	"time"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

type (
	user struct {
		Id string `json:"id"`
	}

	workspace struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	project struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	timeEntriesInterval struct {
		Start time.Time `json:"start"`
		End   time.Time `json:"end"`
	}

	timeEntries struct {
		Id           string              `json:"id"`
		ProjectId    string              `json:"projectId"`
		WorkspaceId  string              `json:"workspaceId"`
		Description  string              `json:"description"`
		TimeInterval timeEntriesInterval `json:"timeInterval"`
	}

	createTimeEntries struct {
		Start       time.Time `json:"start"`
		End         time.Time `json:"end"`
		Description string    `json:"description"`
		ProjectId   string    `json:"projectId"`
	}
)

func (w *workspace) ToGeneric() *trackers.Workspace {
	return &trackers.Workspace{
		Id:   w.Id,
		Name: w.Name,
	}
}

func (p *project) ToGeneric() *trackers.Project {
	return &trackers.Project{
		Id:   p.Id,
		Name: p.Name,
	}
}

func (te *timeEntries) ToGeneric() *trackers.TimeEntries {
	return &trackers.TimeEntries{
		Id:          te.Id,
		ProjectId:   te.ProjectId,
		Start:       te.TimeInterval.Start,
		End:         te.TimeInterval.End,
		Description: te.Description,
	}
}
