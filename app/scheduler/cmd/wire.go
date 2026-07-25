//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"log/slog"
	"scheduler/internal/biz"
	"scheduler/internal/config"
	"scheduler/internal/data"
	"scheduler/internal/server"
	"scheduler/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(
	*config.Bootstrap,
	*common.Server,
	*slog.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
		newApp,
	))
}
