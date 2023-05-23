package service

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/models"
	"unicode/utf8"

	"github.com/go-kratos/kratos/v2/log"
)

type NaviService struct {
	pb.UnimplementedNaviServer
	uc   *biz.NaviUsecase
	xlog *log.Helper
}

func NewNaviService(uc *biz.NaviUsecase, logger log.Logger) *NaviService {
	return &NaviService{
		uc:   uc,
		xlog: log.NewHelper(logger),
	}
}

func (s *NaviService) CreateNavi(ctx context.Context, req *pb.CreateNaviRequest) (*pb.CreateNaviReply, error) {
	if req.Name == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.Name) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	res, erk := s.uc.CreateNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.CreateNaviReply{
		Name: res.Name,
	}, nil
}
func (s *NaviService) DeleteNavi(ctx context.Context, req *pb.DeleteNaviRequest) (*pb.DeleteNaviReply, error) {
	if req.Name == "" {
		return nil, pb.ErrorMissingParam("navi name is missing")
	}
	if utf8.RuneCountInString(req.Name) > 100 {
		return nil, pb.ErrorBadParam("navi name is too long")
	}
	erk := s.uc.DeleteNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.DeleteNaviReply{}, nil
}
func (s *NaviService) SortNavi(ctx context.Context, req *pb.SortNaviRequest) (*pb.SortNaviReply, error) {
	if len(req.Names) == 0 {
		return nil, pb.ErrorMissingParam("navi names is missing")
	}
	for _, name := range req.Names {
		if utf8.RuneCountInString(name) > 100 {
			return nil, pb.ErrorBadParam("navi name is too long")
		}
	}
	erk := s.uc.SortNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.SortNaviReply{}, nil
}
func (s *NaviService) ListNavi(ctx context.Context, req *pb.ListNaviRequest) (*pb.ListNaviReply, error) {
	res, erk := s.uc.ListNavi(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.ListNaviReply{Items: navi2pbs(res)}, nil
}

func navi2pbs(a []*models.Navi) []*pb.NaviType {
	var res = make([]*pb.NaviType, 0, len(a))
	for _, x := range a {
		res = append(res, navi2pb(x))
	}
	return res
}

func navi2pb(x *models.Navi) *pb.NaviType {
	return &pb.NaviType{
		Name:  x.Name,
		Index: x.Index,
	}
}

func (s *NaviService) GetGuestSettings(ctx context.Context, req *pb.GetGuestSettingsRequest) (*pb.GetGuestSettingsReply, error) {
	return &pb.GetGuestSettingsReply{}, nil
}
func (s *NaviService) SetNotRemindInfo(ctx context.Context, req *pb.SetNotRemindInfoRequest) (*pb.SetNotRemindInfoReply, error) {
	return &pb.SetNotRemindInfoReply{}, nil
}
