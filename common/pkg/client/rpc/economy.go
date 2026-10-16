package rpc

import (
	"common/pkg/client/localrpc"
	economyv1 "common/proto/gen/economy/v1"

	"google.golang.org/grpc"
)

type EconomyClient struct {
	Economy           economyv1.EconomyServiceClient
	TCCPoints         economyv1.TCCPointsServiceClient
	TCCPointsTransfer economyv1.TCCPointsTransferServiceClient
}

func NewEconomyClient(conn grpc.ClientConnInterface) *EconomyClient {
	return &EconomyClient{
		Economy:           economyv1.NewEconomyServiceClient(conn),
		TCCPoints:         economyv1.NewTCCPointsServiceClient(conn),
		TCCPointsTransfer: economyv1.NewTCCPointsTransferServiceClient(conn),
	}
}

func MountEconomyServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&economyv1.EconomyService_ServiceDesc, service)
		conn.RegisterMatching(&economyv1.TCCPointsService_ServiceDesc, service)
		conn.RegisterMatching(&economyv1.TCCPointsTransferService_ServiceDesc, service)
	}
}
