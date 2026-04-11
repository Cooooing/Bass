package rpc

import (
	connectorv1 "common/api/gen/connector/v1"

	"google.golang.org/grpc"
)

type ConnectorClient struct {
	connectorv1.ConnectorServiceClient
}

func NewConnectorClient(conn *grpc.ClientConn) *ConnectorClient {
	return &ConnectorClient{
		ConnectorServiceClient: connectorv1.NewConnectorServiceClient(conn),
	}
}
