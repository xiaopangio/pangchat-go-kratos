package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	logic "logic/api/v1"
	"logic/internal/biz"
)

type LogicService struct {
	logic.UnimplementedLogicServer
	biz    *biz.LogicBiz
	helper *log.Helper
}

func NewLogicService(biz *biz.LogicBiz, helper *log.Helper) *LogicService {
	return &LogicService{biz: biz, helper: helper}
}

func (s *LogicService) GetConnectorUrl(ctx context.Context, req *logic.GetConnectorUrlRequest) (*logic.GetConnectorUrlResponse, error) {
	s.helper.Info("GetConnectorUrl")
	err, host, port := s.biz.GetConnectorUrl(ctx)
	if err != nil {
		s.helper.Error("GetConnectorUrl", "err", err)
		return nil, err
	}
	return &logic.GetConnectorUrlResponse{
		Host: host,
		Port: port,
	}, nil
}
