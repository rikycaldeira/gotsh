package timeular

import (
	"fmt"
	"time"

	common "github.com/rikycaldeira/gotsh/common"
	gotshHttp "github.com/rikycaldeira/gotsh/http"
)

const TIMEULAR_API = "https://api.timeular.com/api/v3"

func GetMe(token string) Me {
	url := TIMEULAR_API + "/me"
	me := Me{}

	gotshHttp.Get(url, &me, map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})

	return me
}

func GetActivities(token string) Activities {
	url := TIMEULAR_API + "/activities"
	activities := Activities{}

	gotshHttp.Get(url, &activities, map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})

	return activities
}

func GetRawEntries(token string, interval time.Duration) RawTimeEntries {
	currentTime := time.Now()
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 0, 0, 0, 0, currentTime.Location())
	startTime := endTime.AddDate(0, 0, -1*int(interval.Hours()/24))
	url := TIMEULAR_API +
		fmt.Sprintf("/time-entries/%s/%s", startTime.UTC().Format(common.TIMESTAMP_FORMAT), endTime.UTC().Format(common.TIMESTAMP_FORMAT))

	rawTimeEntries := RawTimeEntries{}
	gotshHttp.Get(url, &rawTimeEntries, map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)})

	return rawTimeEntries
}
