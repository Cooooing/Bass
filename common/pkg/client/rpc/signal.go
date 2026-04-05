package rpc

import (
	signalv1 "common/gen/signal/v1"

	"google.golang.org/grpc"
)

type SignalNodeClient struct {
	signalv1.SignalNodeServiceClient
}

func NewSignalNodeClient(conn *grpc.ClientConn) *SignalNodeClient {
	return &SignalNodeClient{
		SignalNodeServiceClient: signalv1.NewSignalNodeServiceClient(conn),
	}
}
