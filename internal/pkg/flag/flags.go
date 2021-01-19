package flag

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

type (
	Config struct {
		ConfigPath CfgPath
	}
)

const (
	AppName = "clockifier"
)
var (
	app = kingpin.New(AppName, "Clockifier is a tool to import time entries into clockify")
	ConfigCmd = app.Command("config", "Manage the cli configuration")
	ConfigInitCmd = ConfigCmd.Command("init", "Configure the cli to connect with Clockify and Toggl")
	ConfigMappingCmd = ConfigCmd.Command("mapping", "Re-configure the project mapping")
)

func Parse() (*Config, string, error) {
	var cfg Config
	app.Version("v0.0.1")
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	app.Flag("flag", "The path to the configuration file").Default("~/.config/clockifier/clockifier.conf").Short('c').SetValue(&cfg.ConfigPath)
	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		return nil, "", err
	}
	return &cfg, cmd, nil
}
