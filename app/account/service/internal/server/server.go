package server

import (
	"github.com/google/wire"
	"github.com/lalifeier/vgo/app/account/service/internal/conf"

	consul "github.com/go-kratos/consul/registry"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewConsulCRegister)

func NewConsulCRegister(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	consulClient, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(consulClient, consul.WithHealthCheck(false))
	return r
}
