package server

import (
	"context"
	"hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/conf"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/service"
	"hui-webpage-navigation/internal/utils/utils_kratos/utils_kratos_account_auth"

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
			utils_kratos_account_auth.NewMiddleware(newCheckAdminOrNoAccConfig(data), logger),
			utils_kratos_account_auth.NewMiddleware(newCheckAnonymousConfig(data), logger),
			utils_kratos_account_auth.NewMiddleware(newCheckAdminConfig(data), logger),
			utils_kratos_account_auth.NewMiddleware(newCheckAdminSortConfig(data), logger),
			utils_kratos_account_auth.NewMiddleware(newCheckAdminEditConfig(data), logger),
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

func newCheckAdminOrNoAccConfig(data *data.Data) *utils_kratos_account_auth.Config {
	return utils_kratos_account_auth.NewConfig("anonymous", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
		"/api.webnavigation.Ping/Ping",
		"/api.webnavigation.Admin/CreateAdmin",
	}, data.CheckAdminOrNoAccToken)
}

func newCheckAnonymousConfig(data *data.Data) *utils_kratos_account_auth.Config {
	return utils_kratos_account_auth.NewConfig("login", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
		"/api.webnavigation.Admin/AdminLogin",
	}, data.CheckAnonymousToken)
}

// 把需要鉴权的操作列在这里
// 开发者失误了没有把不同权限的操作放在不同的地方
func newCheckAdminConfig(data *data.Data) *utils_kratos_account_auth.Config {
	return utils_kratos_account_auth.NewConfig("get_admin", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
		"/api.webnavigation.Admin/GetAdmin",
		"/api.webnavigation.Admin/ListAdmin",
		"/api.webnavigation.Admin/ListAdminOfMine",
	}, data.CheckAdminToken)
}

func newCheckAdminSortConfig(data *data.Data) *utils_kratos_account_auth.Config {
	return utils_kratos_account_auth.NewConfig("sort", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
		"/api.webnavigation.Navi/SortNavi",
		"/api.webnavigation.NaviLvl2/SortNaviLvl2",
	}, func(ctx context.Context, token string) (context.Context, *errors.Error) {
		ctx, erk := data.CheckAdminToken(ctx, token)
		if erk != nil {
			return nil, erk
		}
		admin, erk := biz.GetAdminFromContext(ctx)
		if erk != nil {
			return nil, erk
		}
		if !admin.CanSort {
			return nil, webnavigation.ErrorAdminNoPermission("can not sort")
		}
		return ctx, nil
	})
}

func newCheckAdminEditConfig(data *data.Data) *utils_kratos_account_auth.Config {
	return utils_kratos_account_auth.NewConfig("edit", true, "admin_manage_token", true, []string{ //跳过鉴权的请求
		"/api.webnavigation.Navi/CreateNavi",
		"/api.webnavigation.Navi/DeleteNavi",
		"/api.webnavigation.NaviLvl2/CreateNaviLvl2",
		"/api.webnavigation.NaviLvl2/DeleteNaviLvl2",
	}, func(ctx context.Context, token string) (context.Context, *errors.Error) {
		ctx, erk := data.CheckAdminToken(ctx, token)
		if erk != nil {
			return nil, erk
		}
		admin, erk := biz.GetAdminFromContext(ctx)
		if erk != nil {
			return nil, erk
		}
		if !admin.CanEdit {
			return nil, webnavigation.ErrorAdminNoPermission("can not edit")
		}
		return ctx, nil
	})
}
