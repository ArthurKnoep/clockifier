package app

import (
	"fmt"
	"os"
	"time"

	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/flag"
	"github.com/ArthurKnoep/toggl-to-clockify/internal/pkg/trackers"
)

func (a *App) parseImportFlag(arg *flag.Argument) (time.Time, time.Time) {
	if arg.ImportAt != "" && (arg.ImportFrom != "" || arg.ImportTo != "") {
		a.logger.Error("Cannot use at the same time the flags at and from or to")
		os.Exit(1)
	}
	if arg.ImportFrom == "" && arg.ImportTo == "" {
		if arg.ImportAt != "" {
			date, err := time.Parse("02/01/2006", arg.ImportAt)
			if err != nil {
				a.logger.WithError(err).Error("Unable to parse date")
				os.Exit(1)
			}
			return date.Add(-time.Hour * time.Duration(date.Hour())), date.Add(time.Hour * (24 - time.Duration(date.Hour())))
		} else {
			now := time.Now().Truncate(time.Hour)
			return now.Add(-time.Hour * time.Duration(now.Hour())), now.Add(time.Hour * (24 - time.Duration(now.Hour())))

		}
	}
	var err error
	from := time.Now()
	to := time.Now()
	if arg.ImportFrom != "" {
		from, err = time.Parse("02/01/2006", arg.ImportFrom)
		if err != nil {
			a.logger.WithError(err).Error("Unable to parse data")
			os.Exit(1)
		}
	}
	if arg.ImportTo != "" {
		to, err = time.Parse("02/01/2006", arg.ImportTo)
		if err != nil {
			a.logger.WithError(err).Error("Unable to parse data")
			os.Exit(1)
		}
	}
	return from.Add(-time.Hour * time.Duration(from.Hour())), to.Add(time.Hour * (24 - time.Duration(to.Hour())))
}

func (a *App) listTimeEntries(from, to time.Time, tracker trackers.Trackers) []*trackers.TimeEntries {
	a.loader.Start()
	timeEntries, err := tracker.ListTimeEntries(from, to)
	if err != nil {
		a.loader.Stop()
		a.logger.WithError(err).Errorf("Unable to list time entries from %s", tracker.Name())
		os.Exit(1)
	}
	a.loader.Stop()
	var info = "entry"
	if len(timeEntries) > 1 {
		info = "entries"
	}
	fmt.Printf("[%s] Found %d time %s\n", tracker.Name(), len(timeEntries), info)
	return timeEntries
}

func (a *App) ImportCmd(arg *flag.Argument) {
	a.applyConfig()
	from, to := a.parseImportFlag(arg)
	timeEntries := a.listTimeEntries(from, to, a.toggl)
	//for _, timeEntry := range timeEntries {
	//	fmt.Println(timeEntry.Id)
	//	fmt.Println(timeEntry.ProjectId)
	//	fmt.Println(timeEntry.Start)
	//	fmt.Println(timeEntry.End)
	//	fmt.Println(timeEntry.Description)
	//}
	timeEntries2 := a.listTimeEntries(from, to, a.clockify)
	//for _, timeEntry := range timeEntries2 {
	//	fmt.Println(timeEntry.Id)
	//	fmt.Println(timeEntry.ProjectId)
	//	fmt.Println(timeEntry.Start)
	//	fmt.Println(timeEntry.End)
	//	fmt.Println(timeEntry.Description)
	//}
	_ = timeEntries
	_ = timeEntries2
}
