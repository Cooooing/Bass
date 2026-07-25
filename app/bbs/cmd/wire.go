//go:build wireinject
// +build wireinject

package main

import (
	"bbs/internal/biz"
	"bbs/internal/config"
	"bbs/internal/data"
	"bbs/internal/server"
	"bbs/internal/service"
	"common/proto/gen/common"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(
	*config.Bootstrap,
	*common.Server,
	*slog.Logger,
) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
		newApp,
	))
}
