//go:build wireinject
// +build wireinject

package main

import (
	"bbs_admin/internal/config"
	"bbs_admin/internal/data"
	"bbs_admin/internal/server"
	"bbs_admin/internal/service"
	"common/proto/gen/common"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, service.ServiceProviderSet, data.DataProviderSet, newApp))
}
