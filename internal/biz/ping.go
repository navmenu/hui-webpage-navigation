package biz

import (
	"context"
	"hui-webpage-navigation/internal/data"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// PingUsecase is a Greeter usecase.
type PingUsecase struct {
	repo *data.PingRepo
	log  *log.Helper
}

// NewPingUsecase new a Greeter usecase.
func NewPingUsecase(repo *data.PingRepo, logger log.Logger) *PingUsecase {
	return &PingUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *PingUsecase) Ping(ctx context.Context, req *wrapperspb.StringValue) (string, *errors.Error) {
	return uc.repo.Ping(ctx, req.GetValue())
}
