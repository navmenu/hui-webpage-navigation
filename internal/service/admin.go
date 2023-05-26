package service

import (
	"context"
	pb "hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/models"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type AdminService struct {
	pb.UnimplementedAdminServer
	uc   *biz.AdminUsecase
	xlog *log.Helper
}

func NewAdminService(uc *biz.AdminUsecase, logger log.Logger) *AdminService {
	return &AdminService{
		uc:   uc,
		xlog: log.NewHelper(logger),
	}
}

func (s *AdminService) CreateAdmin(ctx context.Context, req *pb.CreateAdminRequest) (*pb.CreateAdminReply, error) {
	if req.Username == "" {
		return nil, pb.ErrorBadParam("no username")
	}
	if req.Password == "" {
		return nil, pb.ErrorBadParam("no password")
	}
	if req.Nickname == "" {
		return nil, pb.ErrorBadParam("no nickname")
	}
	admin, erk := s.uc.CreateAdmin(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.CreateAdminReply{
		Username: admin.Username,
		Nickname: admin.Nickname,
	}, nil
}
func (s *AdminService) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginReply, error) {
	if req.Username == "" {
		return nil, pb.ErrorBadParam("no username")
	}
	if req.Password == "" {
		return nil, pb.ErrorBadParam("no password")
	}
	admin, erk := s.uc.AdminLogin(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.AdminLoginReply{
		Username:   admin.Username,
		Nickname:   admin.Nickname,
		Token:      admin.Token,
		ExpireAtNs: time.Now().AddDate(100, 0, 0).UnixNano(),
	}, nil
}
func (s *AdminService) UpdateAdmin(ctx context.Context, req *pb.UpdateAdminRequest) (*pb.UpdateAdminReply, error) {
	return &pb.UpdateAdminReply{}, nil
}
func (s *AdminService) DeleteAdmin(ctx context.Context, req *pb.DeleteAdminRequest) (*pb.DeleteAdminReply, error) {
	return &pb.DeleteAdminReply{}, nil
}
func (s *AdminService) GetAdmin(ctx context.Context, req *pb.GetAdminRequest) (*pb.GetAdminReply, error) {
	admin, erk := s.uc.GetAdmin(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.GetAdminReply{
		Username:    admin.Username,
		Nickname:    admin.Nickname,
		Permissions: permissions2map(admin),
	}, nil
}

func permissions2map(admin *models.Admin) map[string]bool {
	return map[string]bool{
		"can_create_admin": admin.CanCreateAdmin,
		"can_select_admin": admin.CanSelectAdmin,
		"can_edit":         admin.CanEdit,
		"can_sort":         admin.CanSort,
	}
}
func (s *AdminService) ListAdmin(ctx context.Context, req *pb.ListAdminRequest) (*pb.ListAdminReply, error) {
	admin, erk := biz.GetAdminFromContext(ctx)
	if erk != nil {
		return nil, erk
	}
	if !admin.CanSelectAdmin {
		return nil, pb.ErrorAdminNoPermission("username=%v can_select_admin=%v", admin.Username, admin.CanSelectAdmin)
	}
	admins, erk := s.uc.ListAdmin(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.ListAdminReply{Items: admins2pbAdminListItems(admins)}, nil
}

func admins2pbAdminListItems(admins []*models.Admin) []*pb.AdminListItem {
	var items = make([]*pb.AdminListItem, 0, len(admins))
	for _, x := range admins {
		items = append(items, &pb.AdminListItem{
			Username:       x.Username,
			Nickname:       x.Nickname,
			Permissions:    permissions2map(x),
			CreatedByUname: x.CreatedByUname,
		})
	}
	return items
}
func (s *AdminService) ListAdminOfMine(ctx context.Context, req *pb.ListAdminRequest) (*pb.ListAdminReply, error) {
	admin, erk := biz.GetAdminFromContext(ctx)
	if erk != nil {
		return nil, erk
	}
	if !admin.CanCreateAdmin {
		return nil, pb.ErrorAdminNoPermission("username=%v can_create_admin=%v", admin.Username, admin.CanCreateAdmin)
	}
	admins, erk := s.uc.ListAdminOfMine(ctx, req)
	if erk != nil {
		return nil, erk
	}
	return &pb.ListAdminReply{Items: admins2pbAdminListItems(admins)}, nil
}
