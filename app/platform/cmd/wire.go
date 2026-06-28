//go:build wireinject
// +build wireinject

// 构建标签确保该注入桩不会进入最终构建。

package main

import (
	"platform/internal/biz"
	"platform/internal/conf"
	"platform/internal/data"
	"platform/internal/server"
	"platform/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
	"log/slog"
)

// wireApp 初始化 Kratos 应用。
func wireApp(*conf.Bootstrap, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		data.DataProviderSet,
		biz.BizProviderSet,
		service.ServiceProviderSet,
		newApp,
	))
}
