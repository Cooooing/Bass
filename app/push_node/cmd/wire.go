//go:build wireinject
// +build wireinject

package main

import (
	"common/proto/gen/common"
	"log/slog"
	"push_node/internal/biz"
	"push_node/internal/config"
	"push_node/internal/data"
	"push_node/internal/server"
	"push_node/internal/service"

	"github.com/go-kratos/kratos/v3"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func wireApp(bootstrap *config.Bootstrap, serverConf *common.Server, logger *slog.Logger, hubConn *grpc.ClientConn, nodeID string) (*kratos.App, func(), error) {
	panic(wire.Build(server.ServerProviderSet, service.ServiceProviderSet, biz.BizProviderSet, data.DataProviderSet, newApp))
}
