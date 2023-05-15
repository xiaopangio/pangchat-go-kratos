package model

type StsResp struct {
	StatusCode int  `json:"statusCode"`
	Body       Body `json:"body"`
}
type Body struct {
	Credentials Credentials `json:"Credentials"`
	RequestId   string      `json:"RequestId"`
}
type Credentials struct {
	AccessKeyId     string `json:"AccessKeyId"`
	AccessKeySecret string `json:"AccessKeySecret"`
	Expiration      string `json:"Expiration"`
	SecurityToken   string `json:"SecurityToken"`
}
