package app

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/copier"

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
	a.loader.Suffix = fmt.Sprintf(" Looking for time entries in %s", tracker.Name())
	defer func() {a.loader.Suffix = ""}()
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

func (a *App) searchTimeEntry(entry *trackers.TimeEntries, entries []*trackers.TimeEntries) bool {
	for _, timeEntry := range entries {
		if timeEntry.Description == entry.Description {
			startDelay := entry.Start.Sub(timeEntry.Start)
			endDelay := entry.End.Sub(timeEntry.End)
			if startDelay.Truncate(time.Minute) == 0 && endDelay.Truncate(time.Minute) == 0 {
				return true
			}
		}
	}
	return false
}

func (a *App) translateProjectId(srcProjectId string) *string  {
	if projectId, ok := a.cfg.ProjectMapping[srcProjectId]; ok {
		return &projectId
	}
	return nil
}

func (a *App) importTimeEntries(toImport []*trackers.TimeEntries, dest trackers.Trackers) {
	a.loader.Start()
	defer func() {a.loader.Suffix = ""}()
	for i, timeEntry := range toImport {
		a.loader.Suffix = fmt.Sprintf(" [%d/%d]", i + 1, len(toImport))
		if _, err := dest.CreateTimeEntry(timeEntry); err != nil {
			a.loader.Stop()
			a.logger.WithError(err).Errorf("Unable to create time entry into %s", dest.Name())
			os.Exit(1)
		}
	}
	a.loader.Stop()
}

func (a *App) ImportCmd(arg *flag.Argument) {
	a.applyConfig()
	from, to := a.parseImportFlag(arg)
	srcTimeEntries := a.listTimeEntries(from, to, a.toggl)
	destTimeEntries := a.listTimeEntries(from, to, a.clockify)
	toImport := make([]*trackers.TimeEntries, 0)
	for _, timeEntry := range srcTimeEntries {
		translatedId := a.translateProjectId(timeEntry.ProjectId)
		if translatedId == nil {
			continue
		}
		if !a.searchTimeEntry(timeEntry, destTimeEntries) {
			cpyTimeEntry := trackers.TimeEntries{}
			if err := copier.Copy(&cpyTimeEntry, timeEntry); err != nil {
				a.logger.WithError(err).Error("Unable to copy time entry")
				os.Exit(1)
			}
			cpyTimeEntry.ProjectId = *translatedId
			toImport = append(toImport, &cpyTimeEntry)
		}
	}
	var info = "entry"
	if len(toImport) > 1 {
		info = "entries"
	}
	fmt.Printf("\nFound %d time %s to import\n", len(toImport), info)
	a.importTimeEntries(toImport, a.clockify)
}
