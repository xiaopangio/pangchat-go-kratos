package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "online/api/v1/online"
	"online/internal/biz"
)

type OnlineService struct {
	pb.UnimplementedOnlineServer
	helper *log.Helper
	biz    *biz.OnlineBiz
}

func NewOnlineService(helper *log.Helper, biz *biz.OnlineBiz) *OnlineService {
	return &OnlineService{UnimplementedOnlineServer: pb.UnimplementedOnlineServer{}, helper: helper, biz: biz}
}

func (s *OnlineService) RegisterDevice(ctx context.Context, req *pb.RegisterDeviceRequest) (*pb.RegisterDeviceResponse, error) {
	err := s.biz.RegisterDevice(ctx, req.Uid, req.DeviceUrl)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterDeviceResponse{}, nil
}
func (s *OnlineService) UnregisterDevice(ctx context.Context, req *pb.UnregisterDeviceRequest) (*pb.UnregisterDeviceResponse, error) {
	err := s.biz.UnregisterDevice(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.UnregisterDeviceResponse{}, nil
}
func (s *OnlineService) GetOnlineDevices(ctx context.Context, req *pb.GetOnlineDevicesRequest) (*pb.GetOnlineDevicesResponse, error) {
	devices, err := s.biz.GetOnlineDevices(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetOnlineDevicesResponse{
		Devices: devices,
	}, nil
}
func (s *OnlineService) GetOnlineDevice(ctx context.Context, req *pb.GetOnlineDeviceRequest) (*pb.GetOnlineDeviceResponse, error) {
	device, err := s.biz.GetOnlineDevice(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	var resp = &pb.GetOnlineDeviceResponse{}
	if device != nil {
		resp = &pb.GetOnlineDeviceResponse{
			DeviceUrl: device.Url,
		}
	}
	return resp, nil
}
