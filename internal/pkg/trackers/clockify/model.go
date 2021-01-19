package clockify

import "github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"

type (
	workspace struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}
)

func (w *workspace) ToGeneric() *trackers.Workspace {
	return &trackers.Workspace{
		Id:   w.Id,
		Name: w.Name,
	}
}
