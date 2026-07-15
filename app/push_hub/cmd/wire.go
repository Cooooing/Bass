//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"log/slog"
	"push_hub/internal/biz"
	"push_hub/internal/config"
	"push_hub/internal/data"
	"push_hub/internal/server"
	"push_hub/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newApp))
}
