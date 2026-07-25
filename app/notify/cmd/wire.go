//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"log/slog"
	"notify/internal/biz"
	"notify/internal/config"
	"notify/internal/data"
	"notify/internal/server"
	"notify/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(
	*config.Bootstrap,
	*common.Server,
	*slog.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newApp))
}
