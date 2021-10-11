// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/lalifeier/vgo/app/account/interface/internal/biz"
	"github.com/lalifeier/vgo/app/account/interface/internal/conf"
	"github.com/lalifeier/vgo/app/account/interface/internal/data"
	"github.com/lalifeier/vgo/app/account/interface/internal/server"
	"github.com/lalifeier/vgo/app/account/interface/internal/service"
	"go.opentelemetry.io/otel/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger, *conf.Registry, trace.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
