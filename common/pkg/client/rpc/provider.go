package rpc

import (
	"common/pkg/client"
	"common/pkg/constant"

	"google.golang.org/grpc"
)

func ProvideUserClient(
	consul *client.ConsulClient,
) (*UserClient, error) {
	return newServiceClient(consul, constant.UserServiceName.String(), NewUserClient)
}

func ProvideContentClient(
	consul *client.ConsulClient,
) (*ContentClient, error) {
	return newServiceClient(consul, constant.ContentServiceName.String(), NewContentClient)
}

func ProvideIMClient(
	consul *client.ConsulClient,
) (*IMClient, error) {
	return newServiceClient(consul, constant.IMServiceName.String(), NewIMClient)
}

func ProvideNotifyClient(
	consul *client.ConsulClient,
) (*NotifyClient, error) {
	return newServiceClient(consul, constant.NotifyServiceName.String(), NewNotifyClient)
}

func ProvidePlatformClient(
	consul *client.ConsulClient,
) (*PlatformClient, error) {
	return newServiceClient(consul, constant.PlatformServiceName.String(), NewPlatformClient)
}

func ProvideSchedulerClient(
	consul *client.ConsulClient,
) (*SchedulerClient, error) {
	return newServiceClient(consul, constant.SchedulerServiceName.String(), NewSchedulerClient)
}

func ProvideGameTownClient(
	consul *client.ConsulClient,
) (*GameTownClient, error) {
	return newServiceClient(consul, constant.GameTownServiceName.String(), NewGameTownClient)
}

func newServiceClient[T any](
	consul *client.ConsulClient,
	service string,
	newFn func(*grpc.ClientConn) T,
) (T, error) {
	conn, err := consul.GetGrpcConn(service)
	if err != nil {
		var zero T
		return zero, err
	}
	return newFn(conn), nil
}
