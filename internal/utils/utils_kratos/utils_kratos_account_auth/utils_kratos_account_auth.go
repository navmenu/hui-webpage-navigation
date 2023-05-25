package utils_kratos_account_auth

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"go.elastic.co/apm/v2"
)

type Config struct {
	enable     bool
	key        string
	include    bool
	operations []string
	check      CheckFunc
}

type CheckFunc func(ctx context.Context, token string) (context.Context, *errors.Error)

func NewConfig(enable bool, key string, include bool, operations []string, check CheckFunc) *Config {
	return &Config{
		enable:     enable,
		key:        key,
		include:    include,
		operations: operations,
		check:      check,
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
	operationsMap := make(map[string]bool, len(cfg.operations))
	for _, v := range cfg.operations {
		operationsMap[v] = true
	}
	return func(ctx context.Context, operation string) bool {
		if !cfg.enable {
			return false
		}
		exist := operationsMap[operation]
		match := exist == cfg.include
		if match {
			xlog.Debugf("operation=%s exist=%v include=%v match=%v must check auth", operation, exist, cfg.include, match)
		} else {
			xlog.Debugf("operation=%s exist=%v include=%v match=%v skip check auth", operation, exist, cfg.include, match)
		}
		return match
	}
}

func middlewareFunc(cfg *Config, logger log.Logger) middleware.Middleware {
	xlog := log.NewHelper(logger)
	xlog.Infof("new check_account_token_pass.middleware enable=%v key=%v operations=%v include=%v", cfg.enable, cfg.key, len(cfg.operations), cfg.include)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if !cfg.enable {
				xlog.Infof("check_account_token_pass: cfg.enable=false anonymous pass")
				return handler(ctx, req)
			}
			if tp, ok := transport.FromServerContext(ctx); ok {
				tx := apm.TransactionFromContext(ctx)
				sp := tx.StartSpan("check_account_token_pass", "auth", nil)
				defer sp.End()

				token := tp.RequestHeader().Get(cfg.key)
				if token == "" {
					return nil, errors.Unauthorized("UNAUTHORIZED", "check_account_token_pass: auth token is missing")
				}
				ctx, erk := cfg.check(ctx, token)
				if erk != nil {
					return nil, erk
				}
				return handler(ctx, req)
			}
			return nil, errors.Unauthorized("UNAUTHORIZED", "check_account_token_pass: wrong context for middleware")
		}
	}
}
