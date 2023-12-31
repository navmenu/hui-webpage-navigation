// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"hui-webpage-navigation/internal/biz"
	"hui-webpage-navigation/internal/conf"
	"hui-webpage-navigation/internal/data"
	"hui-webpage-navigation/internal/server"
	"hui-webpage-navigation/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	pingRepo := data.NewPingRepo(dataData, logger)
	pingUsecase := biz.NewPingUsecase(pingRepo, logger)
	pingService := service.NewPingService(pingUsecase, logger)
	adminUsecase := biz.NewAdminUsecase(dataData, logger)
	adminService := service.NewAdminService(adminUsecase, logger)
	naviUsecase := biz.NewNaviUsecase(dataData, logger)
	naviService := service.NewNaviService(naviUsecase, logger)
	naviLvl2Usecase := biz.NewNaviLvl2Usecase(dataData, logger)
	naviLvl2Service := service.NewNaviLvl2Service(naviLvl2Usecase, logger)
	grpcServer := server.NewGRPCServer(confServer, dataData, pingService, adminService, naviService, naviLvl2Service, logger)
	httpServer := server.NewHTTPServer(confServer, dataData, pingService, adminService, naviService, naviLvl2Service, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
