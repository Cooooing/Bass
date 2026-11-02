//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"badge/internal/biz"
	"badge/internal/config"
	"badge/internal/data"
	"badge/internal/server"
	"badge/internal/service"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newApp))
}
