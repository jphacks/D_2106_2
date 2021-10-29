package utils

import (
	"time"
)

var layout = "2006-01-02T15:04:05+09:00"

func StringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

func TimeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

func UnixToTime(t int64) time.Time {
	return time.Unix(t, 0)
}

func TimeToUnix(t time.Time) int64 {
	return t.Unix()
}
