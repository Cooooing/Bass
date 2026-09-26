//go:build wireinject
// +build wireinject

package main

import (
	"bff_game_idle/internal/biz"
	"bff_game_idle/internal/config"
	"bff_game_idle/internal/data"
	"bff_game_idle/internal/server"
	"bff_game_idle/internal/service"
	commonClient "common/pkg/client"
	"common/proto/gen/common"
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
