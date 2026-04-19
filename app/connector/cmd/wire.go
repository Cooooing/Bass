//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"connector/internal/biz"
	"connector/internal/conf"
	"connector/internal/data"
	"connector/internal/server"
	"connector/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		data.DataProviderSet,
		biz.BizProviderSet,
		service.ServiceProviderSet,
		newApp,
	))
}
