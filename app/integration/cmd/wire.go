//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"integration/internal/conf"
	"integration/internal/data"
	"integration/internal/server"
	"integration/internal/service"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*conf.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, service.ServiceProviderSet, data.DataProviderSet, newApp))
}
