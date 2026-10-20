package server

import (
	commonclient "common/pkg/client"
	commonserver "common/pkg/server"
	"fmt"
	"log/slog"
	"realtime_hub/internal/config"
	"time"

	kratosgrpc "github.com/go-kratos/kratos/v3/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCServer(config *config.Bootstrap, logger *slog.Logger, observer *commonclient.Observer, services []commonserver.Service) *kratosgrpc.Server {
	options := []kratosgrpc.ServerOption{
		kratosgrpc.Middleware(commonserver.RequestLogContextMiddleware(), observer.ServerMiddleware()),
		kratosgrpc.Options(
			grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
				MinTime:             30 * time.Second,
				PermitWithoutStream: false,
			}),
			grpc.KeepaliveParams(keepalive.ServerParameters{
				Time:    30 * time.Second,
				Timeout: 10 * time.Second,
			}),
		),
	}
	if config.Grpc.GetHost() != "" && config.Grpc.GetPort() != 0 {
		options = append(options, kratosgrpc.Address(fmt.Sprintf("%s:%d", config.Grpc.GetHost(), config.Grpc.GetPort())))
	}
	if config.Grpc.GetTimeout() != nil {
		options = append(options, kratosgrpc.Timeout(config.Grpc.GetTimeout().AsDuration()))
	}
	server := kratosgrpc.NewServer(options...)
	for _, service := range services {
		service.RegisterGrpc(server)
	}
	return server
}
