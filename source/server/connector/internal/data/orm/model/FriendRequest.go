package model

type FriendRequest struct {
	RequestId   string `json:"request_id"`
	RequesterId string `json:"requester_id"`
	ReceiverId  string `json:"receiver_id"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nick_name"`
	Desc        string `json:"desc"`
	Status      string `json:"status"`
}
