package rpc

import (
	"common/pkg/client/localrpc"
	economyv1 "common/proto/gen/economy/v1"

	"google.golang.org/grpc"
)

type EconomyClient struct {
	Economy           economyv1.EconomyServiceClient
	PointsTcc         economyv1.PointsTccServiceClient
	PointsTransferTcc economyv1.PointsTransferTccServiceClient
}

func NewEconomyClient(conn grpc.ClientConnInterface) *EconomyClient {
	return &EconomyClient{
		Economy:           economyv1.NewEconomyServiceClient(conn),
		PointsTcc:         economyv1.NewPointsTccServiceClient(conn),
		PointsTransferTcc: economyv1.NewPointsTransferTccServiceClient(conn),
	}
}

func MountEconomyServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&economyv1.EconomyService_ServiceDesc, service)
		conn.RegisterMatching(&economyv1.PointsTccService_ServiceDesc, service)
		conn.RegisterMatching(&economyv1.PointsTransferTccService_ServiceDesc, service)
	}
}
