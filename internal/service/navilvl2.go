package service

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
)

type NaviLvl2Service struct {
	pb.UnimplementedNaviLvl2Server
}

func NewNaviLvl2Service() *NaviLvl2Service {
	return &NaviLvl2Service{}
}

func (s *NaviLvl2Service) CreateNaviLvl2(ctx context.Context, req *pb.CreateNaviLvl2Request) (*pb.CreateNaviLvl2Reply, error) {
	return &pb.CreateNaviLvl2Reply{}, nil
}
func (s *NaviLvl2Service) DeleteNaviLvl2(ctx context.Context, req *pb.DeleteNaviLvl2Request) (*pb.DeleteNaviLvl2Reply, error) {
	return &pb.DeleteNaviLvl2Reply{}, nil
}
func (s *NaviLvl2Service) SortNaviLvl2(ctx context.Context, req *pb.SortNaviLvl2Request) (*pb.SortNaviLvl2Reply, error) {
	return &pb.SortNaviLvl2Reply{}, nil
}
func (s *NaviLvl2Service) ListNaviLvl2(ctx context.Context, req *pb.ListNaviLvl2Request) (*pb.ListNaviLvl2Reply, error) {
	return &pb.ListNaviLvl2Reply{}, nil
}
