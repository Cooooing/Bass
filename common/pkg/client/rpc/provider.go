package rpc

import (
	"common/pkg/client"
	"common/pkg/constant"

	"google.golang.org/grpc"
)

func ProvideUserClient(consul *client.ConsulClient) (*UserClient, func(), error) {
	return newServiceClient(consul, constant.UserServiceName.String(), NewUserClient)
}

func ProvideContentClient(consul *client.ConsulClient) (*ContentClient, func(), error) {
	return newServiceClient(consul, constant.ContentServiceName.String(), NewContentClient)
}

func ProvideIMClient(consul *client.ConsulClient) (*IMClient, func(), error) {
	return newServiceClient(consul, constant.IMServiceName.String(), NewIMClient)
}

func ProvideInfraClient(consul *client.ConsulClient) (*InfraClient, func(), error) {
	return newServiceClient(consul, constant.InfraServiceName.String(), NewInfraClient)
}

func ProvideNotifyClient(consul *client.ConsulClient) (*NotifyClient, func(), error) {
	return newServiceClient(consul, constant.NotifyServiceName.String(), NewNotifyClient)
}

func ProvideSignalNodeClient(consul *client.ConsulClient) (*SignalNodeClient, func(), error) {
	return newServiceClient(consul, constant.SignalServiceName.String(), NewSignalNodeClient)
}

//func ProvideConnectorClient(consul *client.ConsulClient) (*ConnectorClient, func(), error) {
//	return newServiceClient(consul, constant.ConnectorServiceName.String(), NewConnectorClient)
//}

func newServiceClient[T any](consul *client.ConsulClient, service string, newFn func(*grpc.ClientConn) T) (T, func(), error) {
	conn, err := consul.GetGrpcConn(service)
	if err != nil {
		var zero T
		return zero, nil, err
	}
	return newFn(conn), func() { _ = conn.Close() }, nil
}
