package rpc

import (
	"common/pkg/client"
	"common/pkg/constant"
	commonmodule "common/pkg/module"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

// LocalProviderSet 提供模块化单体模式下的本地 RPC 客户端。
var LocalProviderSet = wire.NewSet(
	ProvideLocalUserClient,
	ProvideLocalContentClient,
	ProvideLocalEconomyClient,
	ProvideLocalIMClient,
	ProvideLocalNotifyClient,
	ProvideLocalPlatformClient,
	ProvideLocalPushHubClient,
	ProvideLocalSchedulerClient,
	ProvideLocalGameTownClient,
)

func ProvideUserClient(consul *client.ConsulClient) (*UserClient, error) {
	return newServiceClient(consul, constant.UserServiceName.String(), NewUserClient)
}

func ProvideContentClient(consul *client.ConsulClient) (*ContentClient, error) {
	return newServiceClient(consul, constant.ContentServiceName.String(), NewContentClient)
}

func ProvideEconomyClient(consul *client.ConsulClient) (*EconomyClient, error) {
	return newServiceClient(consul, constant.EconomyServiceName.String(), NewEconomyClient)
}

func ProvideIMClient(consul *client.ConsulClient) (*IMClient, error) {
	return newServiceClient(consul, constant.IMServiceName.String(), NewIMClient)
}

func ProvideNotifyClient(consul *client.ConsulClient) (*NotifyClient, error) {
	return newServiceClient(consul, constant.NotifyServiceName.String(), NewNotifyClient)
}

func ProvidePlatformClient(consul *client.ConsulClient) (*PlatformClient, error) {
	return newServiceClient(consul, constant.PlatformServiceName.String(), NewPlatformClient)
}

func ProvidePushHubClient(consul *client.ConsulClient) (*PushHubClient, error) {
	return newServiceClient(consul, constant.PushHubServiceName.String(), NewPushHubClient)
}
func ProvideSchedulerClient(consul *client.ConsulClient) (*SchedulerClient, error) {
	return newServiceClient(consul, constant.SchedulerServiceName.String(), NewSchedulerClient)
}

func ProvideGameTownClient(consul *client.ConsulClient) (*GameTownClient, error) {
	return newServiceClient(consul, constant.GameTownServiceName.String(), NewGameTownClient)
}

func newServiceClient[T any](consul *client.ConsulClient, service string, newFn func(grpc.ClientConnInterface) T) (T, error) {
	conn, err := consul.GetGrpcConn(service)
	if err != nil {
		var zero T
		return zero, err
	}
	return newFn(conn), nil
}

func ProvideLocalUserClient(clients *commonmodule.Clients) (*UserClient, error) {
	return commonmodule.Client[*UserClient](clients)
}

func ProvideLocalContentClient(clients *commonmodule.Clients) (*ContentClient, error) {
	return commonmodule.Client[*ContentClient](clients)
}

func ProvideLocalEconomyClient(clients *commonmodule.Clients) (*EconomyClient, error) {
	return commonmodule.Client[*EconomyClient](clients)
}

func ProvideLocalIMClient(clients *commonmodule.Clients) (*IMClient, error) {
	return commonmodule.Client[*IMClient](clients)
}

func ProvideLocalNotifyClient(clients *commonmodule.Clients) (*NotifyClient, error) {
	return commonmodule.Client[*NotifyClient](clients)
}

func ProvideLocalPlatformClient(clients *commonmodule.Clients) (*PlatformClient, error) {
	return commonmodule.Client[*PlatformClient](clients)
}

func ProvideLocalPushHubClient(clients *commonmodule.Clients) (*PushHubClient, error) {
	return commonmodule.Client[*PushHubClient](clients)
}

func ProvideLocalSchedulerClient(clients *commonmodule.Clients) (*SchedulerClient, error) {
	return commonmodule.Client[*SchedulerClient](clients)
}

func ProvideLocalGameTownClient(clients *commonmodule.Clients) (*GameTownClient, error) {
	return commonmodule.Client[*GameTownClient](clients)
}
