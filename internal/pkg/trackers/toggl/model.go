package toggl

import (
	"strconv"
	"time"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

type (
	workspace struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	project struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	timeEntries struct {
		Id          int       `json:"id"`
		ProjectId   int       `json:"pid"`
		WorkspaceId int       `json:"wid"`
		Start       time.Time `json:"start"`
		Stop        time.Time `json:"stop"`
		Description string    `json:"description"`
		Duration    int       `json:"duration"`
	}
)

func (w *workspace) ToGeneric() *trackers.Workspace {
	return &trackers.Workspace{
		Id:   strconv.Itoa(w.Id),
		Name: w.Name,
	}
}

func (p *project) ToGeneric() *trackers.Project {
	return &trackers.Project{
		Id:   strconv.Itoa(p.Id),
		Name: p.Name,
	}
}

func (te *timeEntries) ToGeneric() *trackers.TimeEntries {
	return &trackers.TimeEntries{
		Id:          strconv.Itoa(te.Id),
		ProjectId:   strconv.Itoa(te.ProjectId),
		Start:       te.Start,
		End:         te.Stop,
		Description: te.Description,
	}
}
