package service_online

type UnregisterDeviceRequest struct {
	Uid string `json:"uid" form:"uid" binding:"required" label:"用户ID"`
}
