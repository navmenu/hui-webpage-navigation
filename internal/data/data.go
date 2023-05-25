package data

import (
	"context"
	"hui-webpage-navigation/internal/conf"
	"hui-webpage-navigation/internal/models"
	"hui-webpage-navigation/internal/utils/utils_gorm/utils_gorm_mysql"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPingRepo)

// Data .
type Data struct {
	db   *gorm.DB
	xlog *log.Helper
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	xlog := log.NewHelper(logger)
	if c.Database.Driver != "mysql" {
		panic(c.Database.Driver)
	}
	db, err := utils_gorm_mysql.NewGormDB(c.Database.Source, c.Database.Echo)
	if err != nil {
		panic(err)
	}
	// 自动迁移各种模型
	err = db.AutoMigrate(
		&models.Navi{},
		&models.NaviLvl2{},
		&models.GuestSettings{},
	)
	if err != nil {
		panic(err)
	}
	cleanup := func() {
		xlog.Info("closing the data resources")
		xdb, err := db.DB()
		if err != nil {
			xlog.Errorf("error=%v", err)
		}
		err = xdb.Close()
		if err != nil {
			xlog.Errorf("error=%v", err)
		}
	}
	return &Data{
		db:   db,
		xlog: xlog,
	}, cleanup, nil
}

func (data *Data) DB() *gorm.DB {
	return data.db
}

func (data *Data) CheckManageToken(ctx context.Context, token string) (context.Context, *errors.Error) {
	//TODO 这里需要从DB里根据token/cookie查询管理员账号，但是没搞管理员系统，因此这里硬编码一下，稍后就改
	if token != "123" {
		return nil, errors.Unauthorized("UNAUTHORIZED", "data.CheckManageToken return wrong")
	}
	return ctx, nil
}
