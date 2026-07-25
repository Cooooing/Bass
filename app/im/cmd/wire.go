//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"im/internal/biz"
	"im/internal/config"
	"im/internal/data"
	"im/internal/server"
	"im/internal/service"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(
	*config.Bootstrap,
	*common.Server,
	*slog.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, service.ServiceProviderSet, biz.BizProviderSet, data.DataProviderSet, newApp))
}
