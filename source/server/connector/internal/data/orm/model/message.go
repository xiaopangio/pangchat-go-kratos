package model

type UniversalMessage struct {
	T    string `json:"type"`
	Data any    `json:"data"`
}
