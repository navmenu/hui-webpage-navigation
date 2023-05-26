package utils_kratos_auth_matchs

import (
	"context"

	"github.com/barweiss/go-tuple"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"go.elastic.co/apm/v2"
)

type InputType = tuple.T2[[]string, []CheckFunc]
type CheckFunc func(ctx context.Context, token string) (context.Context, *errors.Error)

type Config struct {
	description   string //这个权限的名称/或者描述
	enable        bool
	key           string
	operationsMap map[string][]CheckFunc
}

func NewConfig(description string, enable bool, key string, checkTuples []*InputType) *Config {
	operationsMap := map[string][]CheckFunc{}
	for _, c := range checkTuples {
		for _, operation := range c.V1 {
			operationsMap[operation] = append(operationsMap[operation], c.V2...)
		}
	}
	return &Config{
		description:   description,
		enable:        enable,
		key:           key,
		operationsMap: operationsMap,
	}
}

func (a *Config) GetEnable() bool {
	if a != nil {
		return a.enable
	}
	return false
}

func (a *Config) GetKey() string {
	if a != nil {
		return a.key
	}
	return ""
}

func NewMiddleware(cfg *Config, logger log.Logger) middleware.Middleware {
	return selector.Server(middlewareFunc(cfg, logger)).Match(matchFunc(cfg, logger)).Build()
}

func matchFunc(cfg *Config, logger log.Logger) selector.MatchFunc {
	xlog := log.NewHelper(logger)
	return func(ctx context.Context, operation string) bool {
		if !cfg.enable {
			return false
		}
		checkFuncs, exist := cfg.operationsMap[operation]
		if exist && len(checkFuncs) > 0 {
			xlog.Debugf("operation=%s exist=%v must check description=%v auth", operation, exist, cfg.description)
		} else {
			xlog.Debugf("operation=%s exist=%v skip check description=%v auth", operation, exist, cfg.description)
		}
		return exist
	}
}

func middlewareFunc(cfg *Config, logger log.Logger) middleware.Middleware {
	xlog := log.NewHelper(logger)
	xlog.Infof("new utils_kratos_auth_matchs.middleware enable=%v key=%v operations_map=%v", cfg.enable, cfg.key, len(cfg.operationsMap))
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if !cfg.enable {
				xlog.Infof("utils_kratos_auth_matchs: cfg.enable=false anonymous pass")
				return handler(ctx, req)
			}
			if tp, ok := transport.FromServerContext(ctx); ok {
				tx := apm.TransactionFromContext(ctx)
				sp := tx.StartSpan("utils_kratos_auth_matchs", "auth", nil)
				defer sp.End()

				token := tp.RequestHeader().Get(cfg.key)
				if token == "" {
					return nil, errors.Unauthorized("UNAUTHORIZED", "utils_kratos_auth_matchs: auth token is missing")
				}
				operation := tp.Operation()
				if operation == "" {
					return nil, errors.Unauthorized("UNAUTHORIZED", "utils_kratos_auth_matchs: no operation")
				}
				checkFuncs := cfg.operationsMap[operation]
				if len(checkFuncs) == 0 {
					return nil, errors.Unauthorized("UNAUTHORIZED", "utils_kratos_auth_matchs: no checkFuncs")
				}
				for _, check := range checkFuncs {
					var erk *errors.Error
					ctx, erk = check(ctx, token)
					if erk != nil {
						return nil, erk
					}
				}
				return handler(ctx, req)
			}
			return nil, errors.Unauthorized("UNAUTHORIZED", "utils_kratos_auth_matchs: wrong context for middleware")
		}
	}
}
