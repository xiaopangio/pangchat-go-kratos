package pkg

import "strconv"

func ParseInt64(n string) int64 {
	res, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		return 0
	}
	return res
}
func FormatInt(n int64) string {
	return strconv.FormatInt(n, 10)
}
