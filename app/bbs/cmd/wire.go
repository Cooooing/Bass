// 构建标签：wireinject。
//go:build wireinject
// +build wireinject

// 构建标签确保该注入桩不会进入最终构建。

package main

import (
	"bbs/internal/biz"
	"bbs/internal/conf"
	"bbs/internal/data"
	"bbs/internal/server"
	"bbs/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
	"log/slog"
)

// wireApp 初始化 Kratos 应用。
func wireApp(*conf.Bootstrap, *slog.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
		newApp,
	))
}
