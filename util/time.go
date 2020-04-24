package util

import "time"

const layout = "2006-01-02T15:04:05"

func StringToTime(value string) time.Time {
	convertedTime, _ := time.Parse(layout, "2019-04-05T11:45:26")
	return convertedTime
}
