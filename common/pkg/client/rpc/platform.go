package rpc

import (
	"common/pkg/client/localrpc"
	platformv1 "common/proto/gen/platform/v1"

	"google.golang.org/grpc"
)

type PlatformClient struct {
	IpResolution platformv1.PlatformIpResolutionServiceClient
	Oss          platformv1.PlatformOssServiceClient
}

func NewPlatformClient(
	conn grpc.ClientConnInterface,
) *PlatformClient {
	return &PlatformClient{
		IpResolution: platformv1.NewPlatformIpResolutionServiceClient(conn),
		Oss:          platformv1.NewPlatformOssServiceClient(conn),
	}
}

func NewLocalPlatformClient[T any](services []T) *PlatformClient {
	conn := localrpc.NewConn()
	MountPlatformServices(conn, services)
	return NewPlatformClient(conn)
}

func MountPlatformServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&platformv1.PlatformIpResolutionService_ServiceDesc, service)
		conn.RegisterMatching(&platformv1.PlatformOssService_ServiceDesc, service)
	}
}
