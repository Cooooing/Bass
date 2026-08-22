//go:build wireinject
// +build wireinject

package module

import (
	"common/pkg/client/rpc"
	"log/slog"

	"github.com/google/wire"
)

func New(*Config, *slog.Logger, *rpc.NotifyClient, *rpc.PlatformClient, *rpc.SchedulerClient) (*Module, func(), error) {
	panic(wire.Build(ProviderSet))
}
