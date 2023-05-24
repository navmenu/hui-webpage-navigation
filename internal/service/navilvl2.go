package service

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/models"
	"unicode/utf8"

	"github.com/go-kratos/kratos/v2/log"
)

type NaviLvl2Service struct {
	pb.UnimplementedNaviLvl2Server
	uc   *biz.NaviLvl2Usecase
	xlog *log.Helper
}

func NewNaviLvl2Service(uc *biz.NaviLvl2Usecase, logger log.Logger) *NaviLvl2Service {
	return &NaviLvl2Service{
		uc:   uc,
		xlog: log.NewHelper(logger),
	}
}

func (s *NaviLvl2Service) CreateNaviLvl2(ctx context.Context, req *pb.CreateNaviLvl2Request) (*pb.CreateNaviLvl2Reply, error) {
	if req.Name == "" {
		return nil, pb.ErrorMissingParam("navi lvl2 name is missing")
	}
	if utf8.RuneCountInString(req.Name) > 100 {
		return nil, pb.ErrorBadParam("navi lvl2 name is too long")
	}
	if req.NaviName == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.NaviName) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	if req.Text == "" {
		return nil, pb.ErrorMissingParam("text is missing")
	}
	if req.Link == "" {
		return nil, pb.ErrorMissingParam("link is missing")
	}
	res, erk := s.uc.CreateNaviLvl2(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.CreateNaviLvl2Reply{
		NaviName: res.NaviName,
		Name:     res.Name,
	}, nil
}
func (s *NaviLvl2Service) DeleteNaviLvl2(ctx context.Context, req *pb.DeleteNaviLvl2Request) (*pb.DeleteNaviLvl2Reply, error) {
	if req.Name == "" {
		return nil, pb.ErrorMissingParam("navi lvl2 name is missing")
	}
	if utf8.RuneCountInString(req.Name) > 100 {
		return nil, pb.ErrorBadParam("navi lvl2 name is too long")
	}
	if req.NaviName == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.NaviName) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	erk := s.uc.DeleteNaviLvl2(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.DeleteNaviLvl2Reply{}, nil
}
func (s *NaviLvl2Service) SortNaviLvl2(ctx context.Context, req *pb.SortNaviLvl2Request) (*pb.SortNaviLvl2Reply, error) {
	if req.NaviName == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.NaviName) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	if len(req.Items) == 0 {
		return nil, pb.ErrorMissingParam("navi lvl2 items is missing")
	}
	for _, item := range req.Items {
		if utf8.RuneCountInString(item.Name) > 100 {
			return nil, pb.ErrorBadParam("navi lvl2 name is too long")
		}
		if item.Sort < 0 {
			return nil, pb.ErrorBadParam("navi lvl2 sort is negative")
		}
	}
	erk := s.uc.SortNaviLvl2(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.SortNaviLvl2Reply{}, nil
}
func (s *NaviLvl2Service) ListNaviLvl2(ctx context.Context, req *pb.ListNaviLvl2Request) (*pb.ListNaviLvl2Reply, error) {
	if req.NaviName == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.NaviName) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	res, erk := s.uc.ListNaviLvl2(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.ListNaviLvl2Reply{Items: naviLvl22pbs(res)}, nil
}

func naviLvl22pbs(a []*models.NaviLvl2) []*pb.NaviLvl2Type {
	var res = make([]*pb.NaviLvl2Type, 0, len(a))
	for _, x := range a {
		res = append(res, naviLvl22pb(x))
	}
	return res
}

func naviLvl22pb(x *models.NaviLvl2) *pb.NaviLvl2Type {
	return &pb.NaviLvl2Type{
		NaviName: x.NaviName,
		Name:     x.Name,
		Text:     x.Text,
		Link:     x.Link,
		IsEscrow: x.IsEscrow,
		Sort:     x.Sort,
	}
}
