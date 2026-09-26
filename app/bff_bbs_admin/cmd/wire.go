//go:build wireinject
// +build wireinject

package main

import (
	"bff_bbs_admin/internal/config"
	"bff_bbs_admin/internal/data"
	"bff_bbs_admin/internal/server"
	"bff_bbs_admin/internal/service"
	"common/proto/gen/common"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, service.ServiceProviderSet, data.DataProviderSet, newApp))
}
