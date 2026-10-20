//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
	"log/slog"
	"realtime_node/internal/biz"
	"realtime_node/internal/config"
	"realtime_node/internal/data"
	"realtime_node/internal/server"
	"realtime_node/internal/service"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, service.ServiceProviderSet, biz.BizProviderSet, data.DataProviderSet, newApp))
}
