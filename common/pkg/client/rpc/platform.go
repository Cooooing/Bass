package rpc

import (
	"common/pkg/client/localrpc"
	platformv1 "common/proto/gen/platform/v1"

	"google.golang.org/grpc"
)

type PlatformClient struct {
	IPResolution platformv1.PlatformIPResolutionServiceClient
	Asset        platformv1.PlatformAssetServiceClient
	Holiday      platformv1.PlatformHolidayServiceClient
}

func NewPlatformClient(
	conn grpc.ClientConnInterface,
) *PlatformClient {
	return &PlatformClient{
		IPResolution: platformv1.NewPlatformIPResolutionServiceClient(conn),
		Asset:        platformv1.NewPlatformAssetServiceClient(conn),
		Holiday:      platformv1.NewPlatformHolidayServiceClient(conn),
	}
}

func MountPlatformServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&platformv1.PlatformIPResolutionService_ServiceDesc, service)
		conn.RegisterMatching(&platformv1.PlatformAssetService_ServiceDesc, service)
		conn.RegisterMatching(&platformv1.PlatformHolidayService_ServiceDesc, service)
	}
}
