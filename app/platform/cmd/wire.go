//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"log/slog"
	"platform/internal/biz"
	"platform/internal/config"
	"platform/internal/data"
	"platform/internal/server"
	"platform/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newApp))
}
