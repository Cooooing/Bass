package rpc

import (
	economyv1 "common/proto/gen/economy/v1"

	"google.golang.org/grpc"
)

type EconomyClient struct {
	Economy economyv1.EconomyServiceClient
}

func NewEconomyClient(conn *grpc.ClientConn) *EconomyClient {
	return &EconomyClient{Economy: economyv1.NewEconomyServiceClient(conn)}
}
