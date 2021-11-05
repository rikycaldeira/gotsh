package main

import (
	"flag"
	"time"

	"github.com/rikycaldeira/gotsh/common"
	"github.com/rikycaldeira/gotsh/jira"
	"github.com/rikycaldeira/gotsh/timeular"
)

func main() {
	username := flag.String("username", "", "the JIRA username")
	password := flag.String("password", "", "the JIRA password")
	timeularApiKey := flag.String("api-key", "", "the Timeular API key")
	timeularApiSecret := flag.String("api-secret", "", "the Timeular API secret")
	confirmationDisabled := flag.Bool("y", false, "disables user confirmation")
	days := flag.Int("days", 1, "the number of days in the past to submit, including the current day")

	flag.Parse()

	authData := timeular.Login(*timeularApiKey, *timeularApiSecret)
	defer timeular.Logout(authData.Token)

	var duration time.Duration = (time.Duration)(*days*common.HOURS_IN_DAY) * time.Hour
	activities := timeular.GetActivities(authData.Token)
	rawTimeEntries := timeular.GetRawEntries(authData.Token, duration)
	timesheetEntries := timeular.ProcessTimeEntries(rawTimeEntries, activities)

	jira.SubmitTimesheet(*username, *password, timesheetEntries, !(*confirmationDisabled))
}
