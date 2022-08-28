// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/biz"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/conf"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/data"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/server"
	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, auth *conf.Auth, registry *conf.Registry, logger log.Logger, tracerProvider *trace.TracerProvider) (*kratos.App, func(), error) {
	client := data.NewDaTaoKeClient(confData, logger)
	taoBaoUseCase := biz.NewTaoBaoUseCase(logger, client)
	taoKeService := service.NewTaoKeService(logger, taoBaoUseCase)
	httpServer := server.NewHTTPServer(confServer, auth, logger, taoKeService)
	grpcServer := server.NewGRPCServer(confServer, auth, logger, taoKeService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
	}, nil
}
