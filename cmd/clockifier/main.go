package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/app"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/flag"
)

func main() {
	logger := logrus.New()
	cfg, arg, cmd, err := flag.Parse()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	application, err := app.New(cfg, logger)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	switch cmd {
	case flag.ConfigInitCmd.FullCommand():
		application.ConfigInitCmd()
	case flag.ConfigMappingCmd.FullCommand():
		application.ConfigMappingCmd()
	case flag.ImportCmd.FullCommand():
		application.ImportCmd(arg)
	}
}
