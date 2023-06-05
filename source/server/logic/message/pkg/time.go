package pkg

import (
	"strconv"
	"time"
)

func CompareTime(t string) bool {
	expire, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		return false
	}
	now := time.Now().Unix()
	return now-expire > 0
}
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
