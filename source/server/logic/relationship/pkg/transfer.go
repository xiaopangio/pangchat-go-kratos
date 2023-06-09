package pkg

import "strconv"

func TransferIntToString(value any) string {
	switch value.(type) {
	case int:
		return strconv.Itoa(value.(int))
	case int8:
		return strconv.Itoa(int(value.(int8)))
	case int16:
		return strconv.Itoa(int(value.(int16)))
	case int32:
		return strconv.Itoa(int(value.(int32)))
	case int64:
		return strconv.Itoa(int(value.(int64)))
	case uint:
		return strconv.Itoa(int(value.(uint)))
	case uint8:
		return strconv.Itoa(int(value.(uint8)))
	case uint16:
		return strconv.Itoa(int(value.(uint16)))
	case uint32:
		return strconv.Itoa(int(value.(uint32)))
	case uint64:
		return strconv.Itoa(int(value.(uint64)))
	default:
		return ""
	}
}
