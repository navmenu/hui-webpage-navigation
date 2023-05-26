package data

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
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
	cfg  *conf.Data
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
		&models.Admin{},
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
		cfg:  c,
		db:   db,
		xlog: xlog,
	}, cleanup, nil
}

func (data *Data) DB() *gorm.DB {
	return data.db
}

func (data *Data) CheckAdminOrNoAccToken(ctx context.Context, token string) (context.Context, *errors.Error) {
	db := data.DB()
	var count int64
	if err := db.WithContext(ctx).Table(models.AdminTable).Count(&count).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	if count != 0 {
		return data.CheckAdminToken(ctx, token)
	} else {
		return data.CheckAnonymousToken(ctx, token)
	}
}

func (data *Data) CheckAnonymousToken(ctx context.Context, token string) (context.Context, *errors.Error) {
	if token != data.cfg.AnonymousToken {
		return nil, pb.ErrorAdminNotCreated("no admin")
	}
	return ctx, nil
}

func (data *Data) CheckAdminToken(ctx context.Context, token string) (context.Context, *errors.Error) {
	db := data.DB()
	var admin models.Admin
	if err := db.WithContext(ctx).Table(models.AdminTable).Where("token=?", token).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pb.ErrorAdminNotFound("error=%v", err)
		}
		return nil, pb.ErrorDbError("error=%v", err)
	}
	ctx = context.WithValue(ctx, models.Admin{}, &admin)
	return ctx, nil
}
