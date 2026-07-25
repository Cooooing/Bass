//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"game_town/internal/biz"
	"game_town/internal/config"
	"game_town/internal/data"
	"game_town/internal/server"
	"game_town/internal/service"
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
