//go:build wireinject
// +build wireinject

package main

import (
	"bff_bbs/internal/biz"
	"bff_bbs/internal/config"
	"bff_bbs/internal/data"
	"bff_bbs/internal/server"
	"bff_bbs/internal/service"
	"common/proto/gen/common"
	"log/slog"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
)

func wireApp(*config.Bootstrap, *common.Server, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
		newApp,
	))
}
