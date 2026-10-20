//go:build wireinject

package main

import (
	"common/proto/gen/common"
	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
	"log/slog"
	"realtime_hub/internal/biz"
	"realtime_hub/internal/config"
	"realtime_hub/internal/data"
	"realtime_hub/internal/server"
	"realtime_hub/internal/service"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, service.ServiceProviderSet, biz.BizProviderSet, data.DataProviderSet, newApp))
}
