package trackers

import "time"

type (
	Workspace struct {
		Id   string
		Name string
	}

	Project struct {
		Id   string
		Name string
	}

	TimeEntries struct {
		Id          string
		ProjectId   string
		Start       time.Time
		End         time.Time
		Description string
		TaskId      *string
	}

	Trackers interface {
		// Name returns the name of the time tracker
		Name() string

		// Test will check if the tracker works (correct API Key, service online, ...)
		Test() error

		// HasWorkspace returns true if the tracker has a notion of workspace
		HasWorkspace() bool
		// ListWorkspaces list the current workspaces of the user
		ListWorkspaces() ([]*Workspace, error)

		// ListProjects list the current project of the user
		ListProjects() ([]*Project, error)

		// ListTimeEntries list time entries between two dates
		ListTimeEntries(from, to time.Time) ([]*TimeEntries, error)
		// CreateTimeEntry create a new time entries
		CreateTimeEntry(entries *TimeEntries) (*TimeEntries, error)
	}
)
