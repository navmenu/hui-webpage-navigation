package server

import (
	"context"
	"hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/conf"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/models"
	"hui-webpage-navigation/internal/service"
	"hui-webpage-navigation/internal/utils/utils_kratos/utils_kratos_auth_matchs"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *conf.Server,
	data *data.Data,
	pingService *service.PingService,
	adminService *service.AdminService,
	naviService *service.NaviService,
	naviLvl2Service *service.NaviLvl2Service,
	logger log.Logger,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			//utils_kratos_auth_match.NewMiddleware(newCheckAdminOrNoAccConfig(data), logger),
			//utils_kratos_auth_match.NewMiddleware(newCheckAnonymousConfig(data), logger),
			//utils_kratos_auth_match.NewMiddleware(newCheckAdminConfig(data), logger),
			//utils_kratos_auth_match.NewMiddleware(newCheckAdminSortConfig(data), logger),
			//utils_kratos_auth_match.NewMiddleware(newCheckAdminEditConfig(data), logger),
			utils_kratos_auth_matchs.NewMiddleware(newCheckConfig(data), logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	webnavigation.RegisterPingServer(srv, pingService)
	webnavigation.RegisterAdminServer(srv, adminService)
	webnavigation.RegisterNaviServer(srv, naviService)
	webnavigation.RegisterNaviLvl2Server(srv, naviLvl2Service)
	return srv
}

func newCheckConfig(data *data.Data) *utils_kratos_auth_matchs.Config {
	return utils_kratos_auth_matchs.NewConfig("check_auth", true, "admin_manage_token", []*utils_kratos_auth_matchs.InputType{
		{
			V1: []string{
				"/api.webnavigation.Ping/Ping",
				"/api.webnavigation.Admin/CreateAdmin",
			},
			V2: []utils_kratos_auth_matchs.CheckFunc{
				data.CheckAdminOrNoAccToken,
			},
		},
		{
			V1: []string{
				"/api.webnavigation.Admin/AdminLogin",
			},
			V2: []utils_kratos_auth_matchs.CheckFunc{
				data.CheckAnonymousToken,
			},
		},
		{
			V1: []string{
				"/api.webnavigation.Admin/GetAdmin",
			},
			V2: []utils_kratos_auth_matchs.CheckFunc{
				data.CheckAdminToken,
			},
		},
		{
			V1: []string{
				"/api.webnavigation.Admin/ListAdmin",
			},
			V2: []utils_kratos_auth_matchs.CheckFunc{
				data.CheckAdminToken,
				CheckAdminCanSelectAdmin,
			},
		},
		{
			V1: []string{
				"/api.webnavigation.Admin/ListAdminOfMine",
			},
			V2: []utils_kratos_auth_matchs.CheckFunc{
				data.CheckAdminToken,
				CheckAdminCanCreateAdmin,
			},
		},
		{
			V1: []string{
				"/api.webnavigation.Navi/SortNavi",
				"/api.webnavigation.NaviLvl2/SortNaviLvl2",
			},
			V2: []utils_kratos_auth_matchs.CheckFunc{
				data.CheckAdminToken,
				CheckAdminCanSort,
			},
		},
		{
			V1: []string{
				"/api.webnavigation.Navi/CreateNavi",
				"/api.webnavigation.Navi/DeleteNavi",
				"/api.webnavigation.NaviLvl2/CreateNaviLvl2",
				"/api.webnavigation.NaviLvl2/DeleteNaviLvl2",
			},
			V2: []utils_kratos_auth_matchs.CheckFunc{
				data.CheckAdminToken,
				CheckAdminCanEdit,
			},
		},
	})
}

//func newCheckAdminOrNoAccConfig(data *data.Data) *utils_kratos_auth_match.Config {
//	return utils_kratos_auth_match.NewConfig("anonymous", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
//		"/api.webnavigation.Ping/Ping",
//		"/api.webnavigation.Admin/CreateAdmin",
//	}, data.CheckAdminOrNoAccToken)
//}

//func newCheckAnonymousConfig(data *data.Data) *utils_kratos_auth_match.Config {
//	return utils_kratos_auth_match.NewConfig("login", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
//		"/api.webnavigation.Admin/AdminLogin",
//	}, data.CheckAnonymousToken)
//}

//// 把需要鉴权的操作列在这里
//// 开发者失误了没有把不同权限的操作放在不同的地方
//func newCheckAdminConfig(data *data.Data) *utils_kratos_auth_match.Config {
//	return utils_kratos_auth_match.NewConfig("get_admin", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
//		"/api.webnavigation.Admin/GetAdmin",
//		"/api.webnavigation.Admin/ListAdmin",
//		"/api.webnavigation.Admin/ListAdminOfMine",
//	}, data.CheckAdminToken)
//}

//func newCheckAdminSortConfig(data *data.Data) *utils_kratos_auth_match.Config {
//	return utils_kratos_auth_match.NewConfig("sort", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
//		"/api.webnavigation.Navi/SortNavi",
//		"/api.webnavigation.NaviLvl2/SortNaviLvl2",
//	}, func(ctx context.Context, token string) (context.Context, *errors.Error) {
//		ctx, erk := data.CheckAdminToken(ctx, token)
//		if erk != nil {
//			return nil, erk
//		}
//		return CheckAdminCanSort(ctx, token)
//	})
//}

//func newCheckAdminEditConfig(data *data.Data) *utils_kratos_auth_match.Config {
//	return utils_kratos_auth_match.NewConfig("edit", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
//		"/api.webnavigation.Navi/CreateNavi",
//		"/api.webnavigation.Navi/DeleteNavi",
//		"/api.webnavigation.NaviLvl2/CreateNaviLvl2",
//		"/api.webnavigation.NaviLvl2/DeleteNaviLvl2",
//	}, func(ctx context.Context, token string) (context.Context, *errors.Error) {
//		ctx, erk := data.CheckAdminToken(ctx, token)
//		if erk != nil {
//			return nil, erk
//		}
//		return CheckAdminCanEdit(ctx, token)
//	})
//}

func CheckAdminCanCreateAdmin(ctx context.Context, token string) (context.Context, *errors.Error) {
	return checkAdmin(ctx, func(admin *models.Admin) bool {
		return admin.CanCreateAdmin
	})
}

func CheckAdminCanSelectAdmin(ctx context.Context, token string) (context.Context, *errors.Error) {
	return checkAdmin(ctx, func(admin *models.Admin) bool {
		return admin.CanSelectAdmin
	})
}

func CheckAdminCanSort(ctx context.Context, token string) (context.Context, *errors.Error) {
	return checkAdmin(ctx, func(admin *models.Admin) bool {
		return admin.CanSort
	})
}

func CheckAdminCanEdit(ctx context.Context, token string) (context.Context, *errors.Error) {
	return checkAdmin(ctx, func(admin *models.Admin) bool {
		return admin.CanEdit
	})
}

func checkAdmin(ctx context.Context, checkFunc func(admin *models.Admin) bool) (context.Context, *errors.Error) {
	admin, erk := biz.GetAdminFromContext(ctx)
	if erk != nil {
		return nil, erk
	}
	if !checkFunc(admin) {
		return nil, webnavigation.ErrorAdminNoPermission("can not sort")
	}
	return ctx, nil
}
