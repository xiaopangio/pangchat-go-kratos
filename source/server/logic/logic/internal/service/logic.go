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
func (s *LogicService) GetToolOptions(ctx context.Context, req *logic.GetToolOptionsRequest) (*logic.GetToolOptionsResponse, error) {
	s.helper.Info("GetToolOptions")
	if options, err := s.biz.GetToolOptions(ctx); err != nil {
		return nil, err
	} else {
		return &logic.GetToolOptionsResponse{
			Options: options,
		}, nil
	}
}
func (s *LogicService) GetPreEmojis(ctx context.Context, req *logic.GetPreEmojisRequest) (*logic.GetPreEmojisResponse, error) {
	s.helper.Info("GetPreEmojis")
	if emojis, err := s.biz.GetPreEmojis(ctx); err != nil {
		return nil, err
	} else {
		return &logic.GetPreEmojisResponse{
			Emojis: emojis,
		}, nil
	}
}
func (s *LogicService) UploadFile(stream logic.Logic_UploadFileServer) error {
	err := s.biz.UploadFile(stream)
	if err != nil {
		return err
	}
	return nil
}
func (s *LogicService) DownloadFile(req *logic.DownloadFileRequest, stream logic.Logic_DownloadFileServer) error {
	err := s.biz.DownloadFile(req.FilePath, stream)
	if err != nil {
		return err
	}
	return nil
}
