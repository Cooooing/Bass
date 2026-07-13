package rpc

import (
	platformv1 "common/proto/gen/platform/v1"

	"google.golang.org/grpc"
)

type PlatformClient struct {
	IpResolution platformv1.PlatformIpResolutionServiceClient
	Oss          platformv1.PlatformOssServiceClient
}

func NewPlatformClient(conn *grpc.ClientConn) *PlatformClient {
	return &PlatformClient{
		IpResolution: platformv1.NewPlatformIpResolutionServiceClient(conn),
		Oss:          platformv1.NewPlatformOssServiceClient(conn),
	}
}
