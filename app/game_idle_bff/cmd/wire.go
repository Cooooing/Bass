//go:build wireinject
// +build wireinject

package main

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"game_idle_bff/internal/biz"
	"game_idle_bff/internal/config"
	"game_idle_bff/internal/data"
	"game_idle_bff/internal/server"
	"game_idle_bff/internal/service"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		commonClient.NewObservability,
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
		newApp,
	))
}
