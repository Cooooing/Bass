//go:build wireinject
// +build wireinject

// 构建标签确保该注入桩不会进入最终构建。

package main

import (
	"integration/internal/conf"
	"integration/internal/data"
	"integration/internal/server"
	"integration/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp 初始化 Kratos 应用。
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		data.DataProviderSet,
		service.ServiceProviderSet,
		newApp,
	))
}
