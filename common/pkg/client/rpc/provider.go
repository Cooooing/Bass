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
	ProvideLocalRealtimeHubClient,
	ProvideLocalSchedulerClient,
	ProvideLocalGameIdleClient,
	ProvideLocalBadgeClient,
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

func ProvideRealtimeHubClient(consul *client.ConsulClient) (*RealtimeHubClient, error) {
	return newServiceClient(consul, constant.RealtimeHubServiceName.String(), NewRealtimeHubClient)
}

func ProvideSchedulerClient(consul *client.ConsulClient) (*SchedulerClient, error) {
	return newServiceClient(consul, constant.SchedulerServiceName.String(), NewSchedulerClient)
}

func ProvideGameIdleClient(consul *client.ConsulClient) (*GameIdleClient, error) {
	return newServiceClient(consul, constant.GameIdleServiceName.String(), NewGameIdleClient)
}

func ProvideBadgeClient(consul *client.ConsulClient) (*BadgeClient, error) {
	return newServiceClient(consul, constant.BadgeServiceName.String(), NewBadgeClient)
}

func newServiceClient[T any](consul *client.ConsulClient, service string, newFn func(grpc.ClientConnInterface) T) (T, error) {
	conn, err := consul.GetGrpcConn(service)
	if err != nil {
		var zero T
		return zero, err
	}
	return newFn(conn), nil
}

// 单体模式下按类型从本地客户端集合解析模块客户端。
func ProvideLocalUserClient(clients *commonmodule.Clients) (*UserClient, error) {
	return clients.Client[*UserClient]()
}

func ProvideLocalContentClient(clients *commonmodule.Clients) (*ContentClient, error) {
	return clients.Client[*ContentClient]()
}

func ProvideLocalEconomyClient(clients *commonmodule.Clients) (*EconomyClient, error) {
	return clients.Client[*EconomyClient]()
}

func ProvideLocalIMClient(clients *commonmodule.Clients) (*IMClient, error) {
	return clients.Client[*IMClient]()
}

func ProvideLocalNotifyClient(clients *commonmodule.Clients) (*NotifyClient, error) {
	return clients.Client[*NotifyClient]()
}

func ProvideLocalPlatformClient(clients *commonmodule.Clients) (*PlatformClient, error) {
	return clients.Client[*PlatformClient]()
}

func ProvideLocalRealtimeHubClient(clients *commonmodule.Clients) (*RealtimeHubClient, error) {
	return clients.Client[*RealtimeHubClient]()
}

func ProvideLocalSchedulerClient(clients *commonmodule.Clients) (*SchedulerClient, error) {
	return clients.Client[*SchedulerClient]()
}

func ProvideLocalGameIdleClient(clients *commonmodule.Clients) (*GameIdleClient, error) {
	return clients.Client[*GameIdleClient]()
}

func ProvideLocalBadgeClient(clients *commonmodule.Clients) (*BadgeClient, error) {
	return clients.Client[*BadgeClient]()
}
