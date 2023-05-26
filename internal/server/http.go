package server

import (
	"hui-webpage-navigation/api/webnavigation"
	"hui-webpage-navigation/internal/conf"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/service"
	"hui-webpage-navigation/internal/utils/utils_kratos/utils_kratos_auth_matchs"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	data *data.Data,
	pingService *service.PingService,
	adminService *service.AdminService,
	naviService *service.NaviService,
	naviLvl2Service *service.NaviLvl2Service,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
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
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	webnavigation.RegisterPingHTTPServer(srv, pingService)
	webnavigation.RegisterAdminHTTPServer(srv, adminService)
	webnavigation.RegisterNaviHTTPServer(srv, naviService)
	webnavigation.RegisterNaviLvl2HTTPServer(srv, naviLvl2Service)
	return srv
}
