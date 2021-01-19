package toggl

import (
	"strconv"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

type (
	workspace struct {
		Id   int
		Name string
	}

	project struct {
		Id   int
		Name string
	}
)

func (w *workspace) ToGeneric() *trackers.Workspace {
	return &trackers.Workspace{
		Id:   strconv.Itoa(w.Id),
		Name: w.Name,
	}
}

func (w *project) ToGeneric() *trackers.Project {
	return &trackers.Project{
		Id:   strconv.Itoa(w.Id),
		Name: w.Name,
	}
}
