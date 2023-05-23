package biz

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/models"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type NaviUsecase struct {
	data *data.Data
	xlog *log.Helper
}

func NewNaviUsecase(data *data.Data, logger log.Logger) *NaviUsecase {
	return &NaviUsecase{
		data: data,
		xlog: log.NewHelper(logger),
	}
}

func (uc *NaviUsecase) CreateNavi(ctx context.Context, req *pb.CreateNaviRequest) (*models.Navi, *errors.Error) {
	db := uc.data.DB()
	if err := db.WithContext(ctx).Table(models.NaviTable).Where("name=?", req.Name).First(&models.Navi{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// do nothing
		} else {
			return nil, pb.ErrorDbError("error=%v", err)
		}
	} else {
		return nil, pb.ErrorAlreadyExist("navi name=%v already exist", req.Name)
	}
	var cnt int64
	if err := db.WithContext(ctx).Table(models.NaviTable).Count(&cnt).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	var navi = &models.Navi{
		ID:    0,
		Name:  req.Name,
		Index: int32(cnt),
	}
	if err := db.WithContext(ctx).Table(models.NaviTable).Create(navi).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return navi, nil
}

func (uc *NaviUsecase) DeleteNavi(ctx context.Context, req *pb.DeleteNaviRequest) *errors.Error {
	db := uc.data.DB()
	var navi models.Navi
	if err := db.WithContext(ctx).Table(models.NaviTable).Where("name=?", req.Name).First(&navi).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return pb.ErrorNotExist("navi name=%v not exist", req.Name)
		} else {
			return pb.ErrorDbError("error=%v", err)
		}
	}
	if err := db.WithContext(ctx).Table(models.NaviTable).Where("name=?", req.Name).Delete(&navi).Error; err != nil {
		return pb.ErrorDbError("error=%v", err)
	}
	return nil
}

func (uc *NaviUsecase) SortNavi(ctx context.Context, req *pb.SortNaviRequest) *errors.Error {
	if err := uc.data.DB().WithContext(ctx).Transaction(func(db *gorm.DB) error {
		for idx, name := range req.Names {
			if err := db.WithContext(ctx).Table(models.NaviTable).Where("name=?", name).UpdateColumn("index", idx).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return pb.ErrorDbError("error=%v", err)
	}
	return nil
}

func (uc *NaviUsecase) ListNavi(ctx context.Context, req *pb.ListNaviRequest) ([]*models.Navi, *errors.Error) {
	db := uc.data.DB()
	var navis []*models.Navi
	if err := db.WithContext(ctx).Table(models.NaviTable).Order("navis.index").Find(&navis).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return navis, nil
}
