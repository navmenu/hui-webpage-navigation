package biz

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/models"

	"github.com/barweiss/go-tuple"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var maxIndex int
	mqs := db.WithContext(ctx).Table(models.NaviTable)
	if req.ParentNvid != 0 {
		mqs = mqs.Where("parent_nvid=?", req.ParentNvid)
	}
	if err := mqs.Select("COALESCE(MAX(navis.sort), -1)").Scan(&maxIndex).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	var navi = &models.Navi{
		ID:         0,
		Name:       req.Name,
		Sort:       int32(maxIndex + 1),
		ParentNvid: req.ParentNvid,
	}
	var erk *errors.Error
	if err := db.Transaction(func(tx *gorm.DB) error {
		if navi.ParentNvid != 0 {
			var parentNavi models.Navi
			if err := db.WithContext(ctx).Table(models.NaviTable).
				Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("id=?", navi.ParentNvid).First(&parentNavi).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					erk = pb.ErrorNotExist("navi id=%v not exist", navi.ParentNvid)
					return erk
				} else {
					erk = pb.ErrorDbError("error=%v", err)
					return erk
				}
			}
		}
		if err := db.WithContext(ctx).Table(models.NaviTable).Create(navi).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		if erk != nil {
			return nil, erk
		}
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return navi, nil
}

func (uc *NaviUsecase) DeleteNavi(ctx context.Context, req *pb.DeleteNaviRequest) *errors.Error {
	var erk *errors.Error
	if err := uc.data.DB().WithContext(ctx).Transaction(func(db *gorm.DB) error {
		var navi models.Navi
		//删除时需要锁住分类数据
		if err := db.WithContext(ctx).Table(models.NaviTable).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("name=?", req.Name).First(&navi).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				erk = pb.ErrorNotExist("navi name=%v not exist", req.Name)
				return erk
			} else {
				erk = pb.ErrorDbError("error=%v", err)
				return erk
			}
		}
		//检查分类是否底下还有分类
		{
			var cnt int64
			if err := db.WithContext(ctx).Table(models.NaviTable).
				Where("parent_nvid=?", navi.ID).Count(&cnt).Error; err != nil {
				erk = pb.ErrorDbError("error=%v", err)
				return erk
			}
			if cnt > 0 {
				erk = pb.ErrorAlreadyExist("next_navis cnt=%v > 0", cnt)
				return erk
			}
		}
		// 是不是 rm -rf 强制删除，以及删除内容
		if req.Force {
			if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Where("navi_name=?", req.Name).Delete(nil).Error; err != nil {
				erk = pb.ErrorDbError("error=%v", err)
				return erk
			}
		} else {
			var cnt int64 //检查内容的数量
			if err := db.WithContext(ctx).Table(models.NaviLvl2Table).
				Where("navi_name=?", req.Name).Count(&cnt).Error; err != nil {
				erk = pb.ErrorDbError("error=%v", err)
				return erk
			}
			//假如不是强制删除，底下还有内容就报错
			if cnt > 0 {
				erk = pb.ErrorAlreadyExist("not force while lvl2 cnt=%v > 0", cnt)
				return erk
			}
		}
		if err := db.WithContext(ctx).Table(models.NaviTable).Where("name=?", req.Name).Delete(&navi).Error; err != nil {
			erk = pb.ErrorDbError("error=%v", err)
			return erk
		}
		return nil
	}); err != nil {
		if erk != nil {
			return erk
		}
		return pb.ErrorDbError("error=%v", err)
	}
	return erk
}

func (uc *NaviUsecase) SortNavi(ctx context.Context, req *pb.SortNaviRequest) *errors.Error {
	if err := uc.data.DB().WithContext(ctx).Transaction(func(db *gorm.DB) error {
		for _, item := range req.Items {
			if err := db.WithContext(ctx).Table(models.NaviTable).Where("name=?", item.Name).UpdateColumn("sort", item.Sort).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return pb.ErrorDbError("error=%v", err)
	}
	return nil
}

func (uc *NaviUsecase) ListNavi(ctx context.Context, req *pb.ListNaviRequest) ([]*tuple.T2[*models.Navi, []*models.NaviLvl2], *errors.Error) {
	db := uc.data.DB()
	var navis []*models.Navi
	if err := db.WithContext(ctx).Table(models.NaviTable).Order("navis.sort").Find(&navis).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	var naviLvl2s []*models.NaviLvl2
	if err := db.WithContext(ctx).Table(models.NaviLvl2Table).Order("navi_lvl2.sort").Find(&naviLvl2s).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	xmp := make(map[string][]*models.NaviLvl2, len(navis))
	for _, x := range naviLvl2s {
		xmp[x.NaviName] = append(xmp[x.NaviName], x)
	}
	res := make([]*tuple.T2[*models.Navi, []*models.NaviLvl2], 0, len(navis))
	for _, navi := range navis {
		subs := xmp[navi.Name]
		delete(xmp, navi.Name)
		res = append(res, &tuple.T2[*models.Navi, []*models.NaviLvl2]{
			V1: navi,
			V2: subs,
		})
	}
	if len(xmp) > 0 {
		for naviName, sub := range xmp {
			uc.xlog.Warnf("some lvl2 not meet navi. navi_name=%v size=%v", naviName, len(sub))
		}
	}
	return res, nil
}

type GuestSettingPK struct {
	Ip     string
	Cookie string
}

func (uc *NaviUsecase) GetGuestSettings(ctx context.Context, req *GuestSettingPK) (*models.GuestSettings, *errors.Error) {
	db := uc.data.DB()
	var guestSettings models.GuestSettings
	if err := db.WithContext(ctx).Table(models.GuestSettingsTable).Where("ip=? and cookie=?", req.Ip, req.Cookie).First(&guestSettings).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pb.ErrorNotExist("error=%v", err)
		}
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return &guestSettings, nil
}

func (uc *NaviUsecase) SetNotRemindInfo(ctx context.Context, req *GuestSettingPK) *errors.Error {
	db := uc.data.DB()
	var guestSettings models.GuestSettings
	if err := db.WithContext(ctx).Table(models.GuestSettingsTable).Where("ip=? and cookie=?", req.Ip, req.Cookie).First(&guestSettings).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			guestSettings = models.GuestSettings{
				ID:              0,
				Ip:              req.Ip,
				Cookie:          req.Cookie,
				IsNotRemindInfo: false,
			}
		} else {
			return pb.ErrorDbError("error=%v", err)
		}
	}
	guestSettings.IsNotRemindInfo = true
	if err := db.WithContext(ctx).Save(&guestSettings).Error; err != nil {
		return pb.ErrorDbError("error=%v", err)
	}
	return nil
}
