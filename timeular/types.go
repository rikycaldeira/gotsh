package timeular

import "time"

type AuthRequest struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

type Auth struct {
	Token string `json:"token"`
}

type Me struct {
	Data struct {
		UserId         string `json:"userId"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		DefaultSpaceID string `json:"defaultSpaceId"`
	} `json:"data"`
}

type Activities struct {
	Activities         []Activity `json:"activities"`
	InactiveActivities []Activity `json:"inactiveActivities"`
	ArchivedActivities []Activity `json:"archivedActivities"`
}

type Activity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Integration string `json:"integration"`
	SpaceID     string `json:"spaceId"`
	DeviceSide  int    `json:"deviceSide"`
}

type Duration struct {
	StartedAt string `json:"startedAt"`
	StoppedAt string `json:"stoppedAt"`
}

type RawTimeEntries struct {
	Data []struct {
		ID         string   `json:"id"`
		ActivityID string   `json:"activityId"`
		Duration   Duration `json:"duration"`
		Note       struct {
			Text     string   `json:"text"`
			Tags     []string `json:"tags"`
			Mentions []string `json:"mentions"`
		} `json:"note"`
	} `json:"timeEntries"`
}

type TimesheetEntry struct {
	ActivityName string
	Ticket       string
	Date         time.Time
	Duration     time.Duration
	Description  string
}
