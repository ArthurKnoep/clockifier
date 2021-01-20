package app

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/sirupsen/logrus"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/config"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/flag"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers/clockify"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers/toggl"
)

type App struct {
	logger    *logrus.Logger
	flag      *flag.Config
	cfg       *config.File
	cfgLoaded bool
	loader    *spinner.Spinner

	clockify *clockify.Clockify
	toggl    *toggl.Toggl
}

func (a *App) applyConfig() {
	if !a.cfgLoaded {
		a.logger.Error("The configuration is not specified, run clockifier config init before")
		os.Exit(1)
	}
	a.clockify = clockify.New(a.cfg.Clockify.ApiKey)
	a.clockify.SetWorkspaceId(a.cfg.Clockify.WorkspaceId)
	a.toggl = toggl.New(a.cfg.Toggl.ApiKey)
	a.toggl.SetWorkspaceId(a.cfg.Toggl.WorkspaceId)
}

func New(flag *flag.Config, logger *logrus.Logger) (*App, error) {
	app := App{
		logger:    logger,
		flag:      flag,
		cfgLoaded: false,
		loader:    spinner.New(spinner.CharSets[14], 100*time.Millisecond),
	}
	configuration, err := config.LoadConfig(flag.ConfigPath.String())
	if err != nil && err != config.NoConfigPresent {
		return nil, err
	} else if err == nil {
		app.cfg = configuration
		app.cfgLoaded = true
	}
	return &app, nil
}
