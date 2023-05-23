package service

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
)

type NaviService struct {
	pb.UnimplementedNaviServer
}

func NewNaviService() *NaviService {
	return &NaviService{}
}

func (s *NaviService) CreateNavi(ctx context.Context, req *pb.CreateNaviRequest) (*pb.CreateNaviReply, error) {
	return &pb.CreateNaviReply{}, nil
}
func (s *NaviService) DeleteNavi(ctx context.Context, req *pb.DeleteNaviRequest) (*pb.DeleteNaviReply, error) {
	return &pb.DeleteNaviReply{}, nil
}
func (s *NaviService) SortNavi(ctx context.Context, req *pb.SortNaviRequest) (*pb.SortNaviReply, error) {
	return &pb.SortNaviReply{}, nil
}
func (s *NaviService) ListNavi(ctx context.Context, req *pb.ListNaviRequest) (*pb.ListNaviReply, error) {
	return &pb.ListNaviReply{}, nil
}
func (s *NaviService) GetGuestSettings(ctx context.Context, req *pb.GetGuestSettingsRequest) (*pb.GetGuestSettingsReply, error) {
	return &pb.GetGuestSettingsReply{}, nil
}
func (s *NaviService) SetNotRemindInfo(ctx context.Context, req *pb.SetNotRemindInfoRequest) (*pb.SetNotRemindInfoReply, error) {
	return &pb.SetNotRemindInfoReply{}, nil
}
