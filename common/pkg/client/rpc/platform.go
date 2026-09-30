package rpc

import (
	"common/pkg/client/localrpc"
	platformv1 "common/proto/gen/platform/v1"

	"google.golang.org/grpc"
)

type PlatformClient struct {
	IpResolution platformv1.PlatformIpResolutionServiceClient
	Asset        platformv1.PlatformAssetServiceClient
}

func NewPlatformClient(
	conn grpc.ClientConnInterface,
) *PlatformClient {
	return &PlatformClient{
		IpResolution: platformv1.NewPlatformIpResolutionServiceClient(conn),
		Asset:        platformv1.NewPlatformAssetServiceClient(conn),
	}
}

func MountPlatformServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&platformv1.PlatformIpResolutionService_ServiceDesc, service)
		conn.RegisterMatching(&platformv1.PlatformAssetService_ServiceDesc, service)
	}
}
