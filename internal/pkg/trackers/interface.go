package trackers

type (
	Workspace struct {
		Id   string
		Name string
	}

	Project struct {
		Id   string
		Name string
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
	}
)
