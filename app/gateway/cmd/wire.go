// go:build wireinject
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"gateway/internal/biz"
	"gateway/internal/conf"
	"gateway/internal/data"
	"gateway/internal/server"
	"gateway/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger, *log.Helper) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
		newApp,
	))
}
