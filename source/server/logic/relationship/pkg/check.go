package pkg

func CheckRequestStatus(status string) bool {
	if status != Pending && status != Agreed && status != Refused {
		return false
	}
	return true
}
