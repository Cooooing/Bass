package rpc

import (
	"common/pkg/client/localrpc"
	badgev1 "common/proto/gen/badge/v1"
	"google.golang.org/grpc"
)

// BadgeClient 聚合 Badge 服务的内部 gRPC 客户端。
type BadgeClient struct {
	Badge badgev1.BadgeServiceClient
}

func NewBadgeClient(conn grpc.ClientConnInterface) *BadgeClient {
	return &BadgeClient{Badge: badgev1.NewBadgeServiceClient(conn)}
}

func MountBadgeServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&badgev1.BadgeService_ServiceDesc, service)
	}
}
