//go:build wireinject
// +build wireinject

package module

import (
	"common/pkg/client/rpc"

	"github.com/google/wire"
)

func New(*Config, *rpc.UserClient, *rpc.ContentClient, *rpc.EconomyClient, *rpc.NotifyClient, *rpc.PlatformClient) (*Module, error) {
	panic(wire.Build(ProviderSet))
}
