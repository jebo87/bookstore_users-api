package date_utils

import "time"

const (
	apiDateLayout   = "2006-01-02T15:04:05Z"
	apiDateDBLayout = "2006-01-02 15:04:05"
)

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
func GetNowDB() string {
	return GetNow().Format(apiDateDBLayout)
}

func GetNow() time.Time {
	return time.Now().UTC()
}
