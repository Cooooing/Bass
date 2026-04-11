package rpc

import (
	infrav1 "common/api/gen/infra/v1"

	"google.golang.org/grpc"
)

type InfraClient struct {
	Email infrav1.InfraEmailServiceClient
	Sms   infrav1.InfraSmsServiceClient
	Oss   infrav1.InfraOssServiceClient
}

func NewInfraClient(conn *grpc.ClientConn) *InfraClient {
	return &InfraClient{
		Email: infrav1.NewInfraEmailServiceClient(conn),
		Sms:   infrav1.NewInfraSmsServiceClient(conn),
		Oss:   infrav1.NewInfraOssServiceClient(conn),
	}
}
