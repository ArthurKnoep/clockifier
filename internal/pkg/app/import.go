package app

import (
	"fmt"
	"os"
	"time"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/flag"
)

func (a *App) parseImportFlag(arg *flag.Argument) (time.Time, time.Time) {
	if arg.ImportAt != "" && (arg.ImportFrom != "" || arg.ImportTo != "") {
		a.logger.Error("Cannot use at the same time the flags at and from or to")
		os.Exit(1)
	}
	if arg.ImportAt != "" && arg.ImportFrom == "" && arg.ImportTo == "" {
		date, err := time.Parse("02/01/2006", arg.ImportAt)
		if err != nil {
			a.logger.WithError(err).Error("Unable to parse date")
			os.Exit(1)
		}
		fmt.Println(date)
		return date.Add(-time.Hour * time.Duration(date.Hour())), date.Add(time.Hour * (24 - time.Duration(date.Hour())))
	} else if arg.ImportAt == "" {
		now := time.Now().Truncate(time.Hour)
		return now.Add(-time.Hour * time.Duration(now.Hour())), now.Add(time.Hour * (24 - time.Duration(now.Hour())))

	}
	return time.Time{}, time.Time{}
}

func (a *App) ImportCmd(arg *flag.Argument) {
	a.applyConfig()
	fmt.Println(a.parseImportFlag(arg))
}
