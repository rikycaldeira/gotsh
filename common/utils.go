package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func AreSameMinute(time1 time.Time, time2 time.Time) bool {
	return time1.Local().YearDay() == time2.Local().YearDay() &&
		time1.Local().Hour() == time2.Local().Hour() &&
		time1.Local().Minute() == time2.Local().Minute()
}

func UserConfirmation(label string, options []string, confirmationValue string) bool {
	return StringPrompt(label, options) == "yes"
}

func StringPrompt(label string, options []string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprintf(os.Stderr, "%s [%s] ", label, strings.Join(options, "|"))
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
