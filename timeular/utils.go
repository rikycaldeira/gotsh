package timeular

import (
	"fmt"
	"regexp"
	"time"

	"github.com/rikycaldeira/gotsh/common"
)

func ProcessTimeEntries(timeEntries RawTimeEntries, activities Activities) []common.TimesheetEntry {
	timesheetEntries := make([]common.TimesheetEntry, 0)
	regex := regexp.MustCompile(`(?P<Ticket>[A-Z]+-[0-9]+)? *(?P<Description>.*)`)
	for _, entry := range timeEntries.Data {
		rawDescription := entry.Note.Text
		matches := regex.FindStringSubmatch(rawDescription)
		ticket := matches[regex.SubexpIndex("Ticket")]
		description := matches[regex.SubexpIndex("Description")]

		startedAtTime, err := time.Parse(common.TIMESTAMP_FORMAT, entry.Duration.StartedAt)
		if err != nil {
			panic(fmt.Sprintf("Could not process started at date '%s'", entry.Duration.StartedAt))
		}
		stoppedAtTime, err := time.Parse(common.TIMESTAMP_FORMAT, entry.Duration.StoppedAt)
		if err != nil {
			panic(fmt.Sprintf("Could not process stopped at date '%s'", entry.Duration.StoppedAt))
		}

		duration := stoppedAtTime.Sub(startedAtTime)

		var activityName string
		for _, activity := range activities.Activities {
			if activity.Id == entry.ActivityID {
				activityName = activity.Name
				break
			}
		}

		timesheetEntry := common.TimesheetEntry{
			ActivityName: activityName,
			Ticket:       ticket,
			Description:  description,
			Date:         startedAtTime,
			Duration:     duration,
		}

		timesheetEntries = append(timesheetEntries, timesheetEntry)
	}

	return timesheetEntries
}
