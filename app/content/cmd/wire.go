//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"content/internal/biz"
	"content/internal/config"
	"content/internal/data"
	"content/internal/server"
	"content/internal/service"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, data.DataProviderSet, biz.BizProviderSet, service.ServiceProviderSet, newApp))
}
