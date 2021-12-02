package helper

import "time"

const (
	TimeFormatISO8601         = "2006-01-02T15:04:05Z"
	RFC3339FullDate           = "2006-01-02"
	TimeFormatISO8601Extended = "2006-01-02T15:04:05.999999999"
)

func TimeToUnixTime(timeInput string, hasLocation bool) (unixTime time.Time, err error) {
	if hasLocation {
		loc, _ := time.LoadLocation("Asia/Bangkok")
		unixTime, err = time.ParseInLocation(TimeFormatISO8601, timeInput, loc)
	} else {
		unixTime, err = time.Parse(TimeFormatISO8601, timeInput)
	}
	return unixTime, err
}
