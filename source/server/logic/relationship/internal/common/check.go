package common

import "relationship/internal/common/constant"

func CheckRequestStatus(status string) bool {
	if status != constant.Pending && status != constant.Agreed && status != constant.Refused {
		return false
	}
	return true
}
