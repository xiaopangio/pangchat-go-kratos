package model

type SmsCode struct {
	ID         string `json:"id"`
	Phone      string `json:"phone"`       // 手机号
	BizID      string `json:"biz_id"`      // 序列号
	Code       string `json:"code"`        // 验证码
	CreateTime int64  `json:"create_time"` // 创建时间
}
