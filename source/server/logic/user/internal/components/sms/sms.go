package sms

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"time"
	"user/internal/components/redis"
	"user/internal/conf"
	"user/internal/data/orm/model"
	"user/pkg"
)

const SmsCodePrefix = "smscode:"
const SmsPassPrefix = "smspass:"

var Expire = time.Minute * 5

type SmsClient struct {
	client   *dysmsapi.Client
	redisCli *redis.Redis
	smsConf  *conf.Sms
}

func NewSmsClient(cf *conf.Bootstrap, redisCli *redis.Redis) *SmsClient {
	aliyunCf := cf.Aliyun
	c := &openapi.Config{
		AccessKeyId:     tea.String(aliyunCf.AccessKey),
		AccessKeySecret: tea.String(aliyunCf.AccessSecret),
		RegionId:        tea.String(aliyunCf.Sms.RegionId),
	}
	c.Endpoint = tea.String(aliyunCf.Sms.Endpoint)
	client, err := dysmsapi.NewClient(c)
	if err != nil {
		panic(err)
	}
	smsClient := &SmsClient{
		client:   client,
		redisCli: redisCli,
		smsConf:  aliyunCf.Sms,
	}
	return smsClient
}
func (s *SmsClient) CreateSmsRequest(phone, code string) (*dysmsapi.SendSmsRequest, error) {
	jsonCode, err := json.Marshal(map[string]string{
		"code": code,
	})
	if err != nil {
		return nil, pkg.InternalError("json marshal error: %s", err)
	}
	return &dysmsapi.SendSmsRequest{
		SignName:      tea.String(s.smsConf.SignName),
		TemplateCode:  tea.String(s.smsConf.TemplateCode),
		PhoneNumbers:  tea.String(phone),
		TemplateParam: tea.String(string(jsonCode)),
	}, nil
}
func (s *SmsClient) SendSms(request *dysmsapi.SendSmsRequest) (*dysmsapi.SendSmsResponse, error) {
	return s.client.SendSms(request)
}
func (s *SmsClient) VerifySmsCode(phone, code string) error {
	v, err := s.redisCli.Get(SmsCodePrefix + phone)
	if err != nil && err != redis.Nil {
		return pkg.InternalError("redis get error: %s", err)
	}
	if v == "" {
		return pkg.InternalError("验证码已过期")
	}
	smsCode := &model.SmsCode{}
	err = json.Unmarshal([]byte(v), smsCode)
	if err != nil {
		return pkg.InternalError("json unmarshal error: %s", err)
	}
	if code != smsCode.Code {
		return pkg.InvalidArgumentError("验证码错误")
	}
	//删除smscode
	err = s.redisCli.Del(SmsCodePrefix + phone)
	if err != nil {
		return err
	}
	return nil
}
