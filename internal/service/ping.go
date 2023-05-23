package service

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/utils/utils_strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type PingService struct {
	pb.UnimplementedPingServer
	uc   *biz.PingUsecase
	xlog *log.Helper
}

func NewPingService(uc *biz.PingUsecase, logger log.Logger) *PingService {
	return &PingService{uc: uc, xlog: log.NewHelper(logger)}
}

func (s *PingService) Ping(ctx context.Context, req *wrapperspb.StringValue) (*wrapperspb.StringValue, error) {
	s.xlog.Infof("[PING]-[PONG] message=[%s]", utils_strings.SOrX(req.GetValue(), "unknown"))
	msg, erk := s.uc.Ping(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &wrappers.StringValue{Value: msg}, nil
}
