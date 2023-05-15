package service_connector

// LoginRequest 登录请求
type LoginRequest struct {
	T        int    `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
}
