package jira

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/andygrunwald/go-jira"
	"github.com/rikycaldeira/gotsh/common"
)

const JIRA_URL = "https://issues.feedzai.com"

func SubmitTimesheet(username string, password string, timesheetEntries []common.TimesheetEntry, confirmation bool) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}
	client, err := jira.NewClient(tp.Client(), JIRA_URL)
	if err != nil {
		panic(err)
	}

	self, _, err := client.User.GetSelf()
	if err != nil {
		panic(err)
	}

	timeTrackingTasks := getTimeTrackingTasks(client)
	worklogsToSubmit := []jira.WorklogRecord{}

	for _, timesheetEntry := range timesheetEntries {
		ticket := timesheetEntry.Ticket

		if ticket == "" {
			if task, ok := timeTrackingTasks[timesheetEntry.ActivityName]; ok {
				ticket = task.Key
			}
		}

		if ticket != "" {
			existingWorklogs, _, err := client.Issue.GetWorklogs(ticket)
			if err != nil {
				panic(err)
			}

			if checkExistingWorklogBySelf(self, existingWorklogs.Worklogs, timesheetEntry.Date) {
				continue
			}

			startedAt := timesheetEntry.Date
			worklog := jira.WorklogRecord{
				IssueID:          ticket,
				Started:          (*jira.Time)(&startedAt),
				TimeSpentSeconds: int(timesheetEntry.Duration.Seconds()),
				Comment:          timesheetEntry.Description,
			}

			worklogsToSubmit = append(worklogsToSubmit, worklog)
		}
	}

	fmt.Println("Preparing to submit the following worklogs:")
	for i, worklog := range worklogsToSubmit {
		fmt.Printf("%d) %s | %s - %s\n",
			i+1,
			worklog.IssueID,
			(*time.Time)(worklog.Started).Format(common.DATETIME_FORMAT),
			worklog.Comment)
	}

	if !confirmation || common.UserConfirmation("\nShould I continue?", []string{"yes", "no"}, "yes") {
		for _, worklog := range worklogsToSubmit {
			fmt.Println("Submitting worklog: " + worklog.Comment)
			client.Issue.AddWorklogRecord(worklog.IssueID, &worklog)
		}
		fmt.Printf("\nFinished submitting worklogs\n")
	} else {
		fmt.Printf("\nSubmission cancelled by user\n")
	}
}

func checkExistingWorklogBySelf(self *jira.User, worklogs []jira.WorklogRecord, checkTime time.Time) bool {
	for _, existingWorklog := range worklogs {
		if existingWorklog.Author.Name != self.Name {
			continue
		}

		if common.AreSameMinute((time.Time)(*existingWorklog.Started), checkTime) {
			return true
		}
	}

	return false
}

func getTimeTrackingTasks(client *jira.Client) map[string]*jira.Subtasks {
	sprintsInBoard, _, err :=
		client.Board.GetAllSprintsWithOptions(common.CICD_BOARD, &jira.GetAllSprintsOptions{State: "active"})
	if err != nil {
		panic(err)
	}

	var activeSprint int
	for _, sprint := range sprintsInBoard.Values {
		if strings.HasPrefix(sprint.Name, common.CICD_SPRINT_PREFIX) {
			activeSprint = sprint.ID
			break
		}
	}

	issuesInSprint, _, err := client.Sprint.GetIssuesForSprint(activeSprint)
	if err != nil {
		panic(err)
	}

	var timeTrackingIssue jira.Issue
	for _, issue := range issuesInSprint {
		if strings.HasPrefix(issue.Fields.Summary, common.CICD_TIME_TRACKING_TASK) {
			timeTrackingIssue = issue
			break
		}
	}

	timeTrackingTasks := make(map[string]*jira.Subtasks)
	regex := regexp.MustCompile(`(?P<Name>[0-9a-zA-Z ]+) \(.*\)`)
	for _, subtask := range timeTrackingIssue.Fields.Subtasks {
		matches := regex.FindStringSubmatch(subtask.Fields.Summary)
		name := matches[regex.SubexpIndex("Name")]
		if activity, ok := common.ACTIVITY_TRANSLATION[name]; ok {
			timeTrackingTasks[activity] = subtask
		}
	}

	return timeTrackingTasks
}
