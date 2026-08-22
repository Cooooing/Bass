//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"log/slog"
	"monolith/internal/config"
	"monolith/internal/module"
	"monolith/internal/server"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		module.ProviderSet,
		newApp,
	))
}
