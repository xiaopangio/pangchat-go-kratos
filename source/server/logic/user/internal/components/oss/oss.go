package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"user/internal/conf"
)

type OSSClient struct {
	Client    *oss.Client
	Bucket    *oss.Bucket
	STSClient *STSClient
	cf        *conf.Oss
}

func NewOSS(cf *conf.Bootstrap, stsCli *STSClient) *OSSClient {
	aliyunCf := cf.Aliyun
	OssClient := &OSSClient{
		STSClient: stsCli,
		cf:        aliyunCf.Oss,
	}
	OssClient.Reset()
	go OssClient.Watch()
	return OssClient
}
func (o *OSSClient) Reset() {
	stsResp := o.STSClient.Serve()
	client, err := oss.New(
		o.cf.Endpoint,
		stsResp.Body.Credentials.AccessKeyId,
		stsResp.Body.Credentials.AccessKeySecret,
		oss.SecurityToken(stsResp.Body.Credentials.SecurityToken),
	)
	if err != nil {
		panic(err)
	}
	o.Client = client
	o.Bucket, err = o.Client.Bucket(o.cf.Bucket)
	if err != nil {
		panic(err)
	}
	log.Info("OSSClient Reset")
}
func (o *OSSClient) Watch() {
	ticker := time.NewTicker(15 * time.Minute)
	for {
		select {
		case <-ticker.C:
			o.Reset()
		}
	}
}
