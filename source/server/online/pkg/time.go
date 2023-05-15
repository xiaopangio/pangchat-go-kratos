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
