package data

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type PingRepo struct {
	data *Data
	log  *log.Helper
}

func NewPingRepo(data *Data, logger log.Logger) *PingRepo {
	return &PingRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (x *PingRepo) Ping(ctx context.Context, message string) (string, *errors.Error) {
	x.log.Debugf("ping message=%s", message)
	err := x.data.db.WithContext(ctx).Exec("SELECT 200").Error
	if err != nil {
		return "", errors.New(500, "DB_ERROR", err.Error())
	}
	return fmt.Sprintf("pong message=%s", message), nil
}
