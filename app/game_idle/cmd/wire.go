//go:build wireinject
// +build wireinject

package main

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"game_idle/internal/biz"
	"game_idle/internal/config"
	"game_idle/internal/data"
	"game_idle/internal/server"
	"game_idle/internal/service"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		commonClient.NewObservability,
		server.ServerProviderSet,
		data.DataProviderSet,
		biz.BizProviderSet,
		service.ServiceProviderSet,
		newApp,
	))
}
