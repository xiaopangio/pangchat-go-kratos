package oss

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"logic/internal/conf"
	"logic/internal/data/orm/model"
)

type STSClient struct {
	*sts20150401.Client
	cf *conf.Aliyun
}

func NewSTS(cf *conf.Bootstrap) *STSClient {
	aliyunCf := cf.Aliyun
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(aliyunCf.AccessKey),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(aliyunCf.AccessSecret),
	}
	// 访问的域名
	config.Endpoint = tea.String(aliyunCf.Sts.Endpoint)
	client, err := sts20150401.NewClient(config)
	if err != nil {
		panic(err)
	}
	return &STSClient{
		Client: client,
		cf:     aliyunCf,
	}
}
func (s *STSClient) Serve() *model.StsResp {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: tea.String(s.cf.AccessKey),
		// 必填，您的 AccessKey Secret
		AccessKeySecret: tea.String(s.cf.AccessSecret),
	}
	// 访问的域名
	config.Endpoint = tea.String(s.cf.Sts.Endpoint)
	client, err := sts20150401.NewClient(config)
	if err != nil {
		panic(err)
	}
	s.Client = client
	// 创建API请求并设置参数
	request := &sts20150401.AssumeRoleRequest{
		RoleArn:         tea.String(s.cf.Sts.RoleArn),
		RoleSessionName: tea.String(s.cf.Sts.RoleSessionName),
	}
	// 发起请求并处理异常
	response, err := s.Client.AssumeRole(request)
	if err != nil {
		panic(err)
	}
	stsResp := &model.StsResp{}
	err = json.Unmarshal([]byte(response.String()), stsResp)
	if err != nil {
		panic(err)
	}
	// 打印返回的body
	return stsResp
}
