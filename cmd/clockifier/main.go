package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/app"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/flag"
)

func main() {
	logger := logrus.New()
	cfg, cmd, err := flag.Parse()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	application, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	switch cmd {
	case flag.ConfigCmd.FullCommand():
		application.ConfigCmd()
	}
}
