package flag

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

type (
	Config struct {
		ConfigPath CfgPath
	}

	Argument struct {
		ImportAt   string
		ImportFrom string
		ImportTo   string
	}
)

const (
	AppName = "clockifier"
)

var (
	app = kingpin.New(AppName, "Clockifier is a tool to import time entries into clockify")

	configCmd        = app.Command("config", "Manage the cli configuration")
	ConfigInitCmd    = configCmd.Command("init", "Configure the cli to connect with Clockify and Toggl")
	ConfigMappingCmd = configCmd.Command("mapping", "Re-configure the project mapping")

	ImportCmd = app.Command("import", "Import your time entries from Toggl to Clockify")
)

func Parse() (*Config, *Argument, string, error) {
	var cfg Config
	var arg Argument
	app.Version("v0.0.1")
	app.VersionFlag.Short('v')
	app.HelpFlag.Short('h')

	app.Flag("flag", "The path to the configuration file").Default("~/.config/clockifier/clockifier.conf").Short('c').SetValue(&cfg.ConfigPath)
	ImportCmd.Flag("at", "The date to look for time entries to import").StringVar(&arg.ImportAt)
	ImportCmd.Flag("from", "A start date to look for time entries to import").StringVar(&arg.ImportFrom)
	ImportCmd.Flag("to", "An end date to look for time entries to import (Default to today)").StringVar(&arg.ImportTo)
	cmd, err := app.Parse(os.Args[1:])
	if err != nil {
		return nil, nil, "", err
	}
	return &cfg, &arg, cmd, nil
}
