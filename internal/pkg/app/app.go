package app

import (
	"github.com/sirupsen/logrus"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/config"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/flag"
)

type App struct {
	logger    *logrus.Logger
	flag      *flag.Config
	cfg       *config.File
	cfgLoaded bool
}

func NewApp(flag *flag.Config, logger *logrus.Logger) (*App, error) {
	app := App{
		logger:    logger,
		flag:      flag,
		cfgLoaded: false,
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
