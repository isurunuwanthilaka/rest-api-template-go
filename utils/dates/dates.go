package dates

import (
	"time"
)

const (
	apiDateFormat = "2006-01-02T15:04:05-0700"
)

func GetNow() time.Time {
	return time.Now()
}

func GetNowString() string {
	return GetNow().Format(apiDateFormat)
}
