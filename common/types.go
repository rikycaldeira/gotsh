package common

import "time"

type Empty struct{}

type TimesheetEntry struct {
	ActivityName string
	Ticket       string
	Date         time.Time
	Duration     time.Duration
	Description  string
}
