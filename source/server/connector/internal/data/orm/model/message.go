package model

import "connector/api/v1/universal"

type UniversalMessage struct {
	T    string `json:"type"`
	Data any    `json:"data"`
}
type UnreadMessageResponse struct {
	List []*universal.UnreadMessageInfo `json:"list"`
}
