package pkg

import "time"

func UnixToTimeSring(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}
