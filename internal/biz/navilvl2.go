package biz

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/models"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NaviLvl2Usecase struct {
	data *data.Data
	xlog *log.Helper
}

func NewNaviLvl2Usecase(data *data.Data, logger log.Logger) *NaviLvl2Usecase {
	return &NaviLvl2Usecase{
		data: data,
		xlog: log.NewHelper(logger),
	}
}

func (uc *NaviLvl2Usecase) CreateNaviLvl2(ctx context.Context, req *pb.CreateNaviLvl2Request) (*models.NaviLvl2, *errors.Error) {
	var naviLvl2 *models.NaviLvl2
	var erk *errors.Error
	if err := uc.data.DB().WithContext(ctx).Transaction(func(db *gorm.DB) error {
		var navi models.Navi
		//添加内容时需要锁住分类数据
		if err := db.WithContext(ctx).Table(models.NaviTable).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("name=?", req.NaviName).First(&navi).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				erk = pb.ErrorNotExist("navi name=%v not exist", req.NaviName)
				return erk
			} else {
				erk = pb.ErrorDbError("error=%v", err)
				return erk
			}
		}
		//检查内容是否已存在
		if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Where("navi_name=? and name=?", navi.Name, req.Name).First(&models.NaviLvl2{}).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				// do nothing
			} else {
				erk = pb.ErrorDbError("error=%v", err)
				return erk
			}
		} else {
			erk = pb.ErrorAlreadyExist("navi lvl2 name=%v already exist", req.Name)
			return erk
		}
		//统计总数以生成排序的序号
		var maxIndex int
		if err := db.WithContext(ctx).Table(models.NaviLvl2Table).
			Where("navi_name=?", navi.Name).
			Select("COALESCE(MAX(navi_lvl2.sort), -1)").Scan(&maxIndex).Error; err != nil {
			erk = pb.ErrorDbError("error=%v", err)
			return erk
		}
		newIndex := maxIndex + 1
		naviLvl2 = &models.NaviLvl2{
			ID:       0,
			NaviName: req.NaviName,
			Name:     req.Name,
			Text:     req.Text,
			Link:     req.Link,
			IsEscrow: req.IsEscrow,
			Sort:     int32(newIndex),
		}
		if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Create(naviLvl2).Error; err != nil {
			erk = pb.ErrorDbError("error=%v", err)
			return erk
		}
		return nil
	}); err != nil {
		if erk != nil {
			return nil, erk
		}
		return nil, pb.ErrorDbError("error=%v", erk)
	}
	return naviLvl2, nil
}

func (uc *NaviLvl2Usecase) DeleteNaviLvl2(ctx context.Context, req *pb.DeleteNaviLvl2Request) *errors.Error {
	db := uc.data.DB()
	var naviLvl2 models.NaviLvl2
	if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Where("navi_name=? and name=?", req.NaviName, req.Name).First(&naviLvl2).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return pb.ErrorNotExist("navi lvl2 name=%v not exist", req.Name)
		} else {
			return pb.ErrorDbError("error=%v", err)
		}
	}
	if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Where("navi_name=? and name=?", req.NaviName, req.Name).Delete(&naviLvl2).Error; err != nil {
		return pb.ErrorDbError("error=%v", err)
	}
	return nil
}

func (uc *NaviLvl2Usecase) SortNaviLvl2(ctx context.Context, req *pb.SortNaviLvl2Request) *errors.Error {
	if err := uc.data.DB().WithContext(ctx).Transaction(func(db *gorm.DB) error {
		for _, item := range req.Items {
			if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Where("navi_name=? and name=?", req.NaviName, item.Name).UpdateColumn("sort", item.Sort).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return pb.ErrorDbError("error=%v", err)
	}
	return nil
}

func (uc *NaviLvl2Usecase) ListNaviLvl2(ctx context.Context, req *pb.ListNaviLvl2Request) ([]*models.NaviLvl2, *errors.Error) {
	db := uc.data.DB()
	var naviLvl2s []*models.NaviLvl2
	if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Where("navi_name=?", req.NaviName).Order("navi_lvl2.sort").Find(&naviLvl2s).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return naviLvl2s, nil
}
