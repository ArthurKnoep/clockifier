package app

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/config"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers/clockify"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers/toggl"
)

func (a *App) askCfgOverride() {
	overrideCfg := &survey.Confirm{
		Message: "A configuration file already exist, do you want to override it ?",
	}
	var confirmConfigEdit bool
	if err := survey.AskOne(overrideCfg, &confirmConfigEdit); err != nil {
		a.logger.Error(err)
		os.Exit(1)
	}
	if !confirmConfigEdit {
		os.Exit(2)
	}
}

func (a *App) askTrackersWorkspaceId(tracker trackers.Trackers) string {
	if tracker.HasWorkspace() {
		workspaces, err := tracker.ListWorkspaces()
		if err != nil {
			a.logger.WithError(err).Errorf("Unable to list %s workspaces", tracker.Name())
			os.Exit(1)
		}
		opts := make([]string, 0, len(workspaces))
		for _, workspace := range workspaces {
			opts = append(opts, workspace.Name)
		}
		workspaceSelector := &survey.Select{
			Message: "Please select a workspace",
			Options: opts,
		}
		var workspaceSelected string
		if err := survey.AskOne(workspaceSelector, &workspaceSelected); err != nil {
			a.logger.Error(err)
			os.Exit(1)
		}
		for _, workspace := range workspaces {
			if workspace.Name == workspaceSelected {
				return workspace.Id
			}
		}
	}
	return ""
}

func (a *App) askClockifyCfg(cfg *config.Clockify) trackers.Trackers {
	clockifyQuestion := &survey.Input{
		Message: "Enter your Clockify API Key:",
		Help: "Visit https://clockify.me/user/settings to obtain your API Key",
	}
	if err := survey.AskOne(clockifyQuestion, &cfg.ApiKey); err != nil {
		a.logger.Error(err)
		os.Exit(1)
	}
	if cfg.ApiKey == "" {
		a.logger.Error("Invalid Clockify API Key")
		os.Exit(1)
	}
	clockifyCom := clockify.New(cfg.ApiKey)
	if err := clockifyCom.Test(); err != nil {
		a.logger.WithError(err).Error("Unable to connect to Clockify, verify your token or your internet connection")
		os.Exit(1)
	}

	cfg.WorkspaceId = a.askTrackersWorkspaceId(clockifyCom)
	clockifyCom.SetWorkspaceId(cfg.WorkspaceId)
	return clockifyCom
}

func (a *App) askTogglCfg(cfg *config.Toggl) trackers.Trackers {
	togglQuestion := &survey.Input{
		Message: "Enter your Toggl API Key:",
		Help: "Visit https://track.toggl.com/profile to obtain your API Key",
	}
	if err := survey.AskOne(togglQuestion, &cfg.ApiKey); err != nil {
		a.logger.Error(err)
		os.Exit(1)
	}
	if cfg.ApiKey == "" {
		a.logger.Error("Invalid Toggl API Key")
		os.Exit(1)
	}
	togglCom := toggl.New(cfg.ApiKey)
	if err := togglCom.Test(); err != nil {
		a.logger.WithError(err).Error("Unable to connect to Toggl, verify your token or your internet connection")
		os.Exit(1)
	}

	cfg.WorkspaceId = a.askTrackersWorkspaceId(togglCom)
	togglCom.SetWorkspaceId(cfg.WorkspaceId)
	return togglCom
}

func (a *App) askProjectMapping(cfg *config.File, source, dest trackers.Trackers) {
	sourceProjects, err := source.ListProjects()
	if err != nil {
		a.logger.WithError(err).Errorf("Unable to load %s projects", source.Name())
		os.Exit(1)
	}
	destProjects, err := dest.ListProjects()
	if err != nil {
		a.logger.WithError(err).Errorf("Unable to load %s projects", source.Name())
		os.Exit(1)
	}
	opts := make([]string, 0, len(destProjects) + 1)
	opts = append(opts, "None")
	for _, destProject := range destProjects {
		opts = append(opts, fmt.Sprintf("(%s) %s", dest.Name(), destProject.Name))
	}
	for _, sourceProject := range sourceProjects {
		projectSelector := &survey.Select{
			Message: fmt.Sprintf("Please select to which project map the project \"(%s) %s\"", source.Name(), sourceProject.Name),
			Options: opts,
		}
		var workspaceSelected string
		if err := survey.AskOne(projectSelector, &workspaceSelected); err != nil {
			a.logger.Error(err)
			os.Exit(1)
		}
	}
}

func (a *App) ConfigInitCmd() {
	var cfg config.File
	if a.cfgLoaded {
		a.askCfgOverride()
	}
	clockifyCom := a.askClockifyCfg(&cfg.Clockify)
	togglCom := a.askTogglCfg(&cfg.Toggl)
	a.askProjectMapping(&cfg, togglCom, clockifyCom)
	if err := config.SaveConfig(a.flag.ConfigPath.String(), &cfg); err != nil {
		a.logger.WithError(err).Error("Unable to save configuration")
		os.Exit(1)
	}
}

func (a *App) ConfigMappingCmd() {
	a.applyConfig()
	a.askProjectMapping(a.cfg, a.toggl, a.clockify)
	if err := config.SaveConfig(a.flag.ConfigPath.String(), a.cfg); err != nil {
		a.logger.WithError(err).Error("Unable to save new configuration")
		os.Exit(1)
	}
}
