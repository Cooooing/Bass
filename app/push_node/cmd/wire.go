//go:build wireinject
// +build wireinject

// 构建标签确保该注入桩不会进入最终构建。

package main

import (
	"push_node/internal/biz"
	"push_node/internal/conf"
	"push_node/internal/data"
	"push_node/internal/server"
	"push_node/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

// wireApp 初始化 Kratos 应用。
// hubConn 和 nodeID 在 main 中提前准备好，通过 wire.Value 注入。
func wireApp(bootstrap *conf.Bootstrap, logger log.Logger, hubConn *grpc.ClientConn, nodeID string) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ServerProviderSet,
		service.ServiceProviderSet,
		biz.BizProviderSet,
		data.DataProviderSet,
		newApp,
	))
}
