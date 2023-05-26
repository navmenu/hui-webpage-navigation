package biz

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/models"
	"hui-webpage-navigation/internal/utils/utils_uuid"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type AdminUsecase struct {
	data *data.Data
	xlog *log.Helper
}

func NewAdminUsecase(data *data.Data, logger log.Logger) *AdminUsecase {
	return &AdminUsecase{
		data: data,
		xlog: log.NewHelper(logger),
	}
}

func (uc *AdminUsecase) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*models.Admin, *errors.Error) {
	db := uc.data.DB()
	var count int64
	if err := db.WithContext(ctx).Table(models.AdminTable).Count(&count).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	var canCreateAdmin = req.CanCreateAdmin
	var canSelectAdmin = req.CanSelectAdmin
	var canEdit = req.CanEdit
	var canSort = req.CanSort
	var createdByUname string
	if count > 0 {
		admin, erk := GetAdminFromContext(ctx)
		if erk != nil {
			return nil, erk
		}
		if !admin.CanCreateAdmin {
			return nil, pb.ErrorAdminNoPermission("can not create new admin")
		}
		createdByUname = admin.Username
		//这里规定新管理员的权限不能高于老管理员
		canCreateAdmin = canCreateAdmin && admin.CanCreateAdmin
		canSelectAdmin = canSelectAdmin && admin.CanSelectAdmin
		canEdit = canEdit && admin.CanEdit
		canSort = canSort && admin.CanSort
	}
	var newAdmin = &models.Admin{
		ID:             0,
		Username:       req.Username,
		Password:       req.Password,
		Nickname:       req.Nickname,
		CanCreateAdmin: canCreateAdmin,
		CanSelectAdmin: canSelectAdmin,
		CanEdit:        canEdit,
		CanSort:        canSort,
		Token:          utils_uuid.NewUUID(), //随机个token避免空值的重复
		CreatedByUname: createdByUname,
	}
	if err := db.WithContext(ctx).Table(models.AdminTable).Create(newAdmin).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return newAdmin, nil
}

func (uc *AdminUsecase) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*models.Admin, *errors.Error) {
	db := uc.data.DB()
	var admin models.Admin
	if err := db.WithContext(ctx).Table(models.AdminTable).
		Where("username=? and password=?", req.Username, req.Password).
		First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, pb.ErrorAdminNotFound("error=%v", err)
		}
		return nil, pb.ErrorDbError("error=%v", err)
	}
	if err := db.WithContext(ctx).Model(&admin).
		Update("token", utils_uuid.NewUUID()).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return &admin, nil
}

func (uc *AdminUsecase) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*models.Admin, *errors.Error) {
	return GetAdminFromContext(ctx)
}

func (uc *AdminUsecase) ListAdmin(ctx context.Context, req *pb.ListAdminRequest) ([]*models.Admin, *errors.Error) {
	db := uc.data.DB()
	var admins []*models.Admin
	if err := db.WithContext(ctx).Table(models.AdminTable).Find(&admins).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return admins, nil
}

func (uc *AdminUsecase) ListAdminOfMine(ctx context.Context, req *pb.ListAdminRequest) ([]*models.Admin, *errors.Error) {
	admin, erk := GetAdminFromContext(ctx)
	if erk != nil {
		return nil, erk
	}
	db := uc.data.DB()
	var admins []*models.Admin
	if err := db.WithContext(ctx).Table(models.AdminTable).Where("created_by_uname=?", admin.Username).Find(&admins).Error; err != nil {
		return nil, pb.ErrorDbError("error=%v", err)
	}
	return admins, nil
}

func GetAdminFromContext(ctx context.Context) (*models.Admin, *errors.Error) {
	value := ctx.Value(models.Admin{})
	if value == nil {
		return nil, pb.ErrorAdminNotFound("no admin in context")
	}
	admin, ok := value.(*models.Admin)
	if !ok {
		return nil, pb.ErrorAdminNotFound("wrong admin in context")
	}
	return admin, nil
}
