// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"prometheus-manager/app/prom_agent/internal/biz"
	"prometheus-manager/app/prom_agent/internal/conf"
	"prometheus-manager/app/prom_agent/internal/data"
	"prometheus-manager/app/prom_agent/internal/server"
	"prometheus-manager/app/prom_agent/internal/service"
	"prometheus-manager/pkg/helper/plog"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(string2 *string) (*kratos.App, func(), error) {
	confBefore := before()
	bootstrap, err := conf.LoadConfig(string2, confBefore)
	if err != nil {
		return nil, nil, err
	}
	confServer := bootstrap.Server
	confData := bootstrap.Data
	log := bootstrap.Log
	logger := plog.NewLogger(log)
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	pingRepo := data.NewPingRepo(dataData, logger)
	pingBiz := biz.NewPingBiz(pingRepo, logger)
	pingService := service.NewPingService(pingBiz, logger)
	grpcServer := server.NewGRPCServer(confServer, pingService, logger)
	httpServer := server.NewHTTPServer(confServer, pingService, logger)
	app := newApp(grpcServer, httpServer, logger)
	return app, func() {
		cleanup()
	}, nil
}
