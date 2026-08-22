//go:build wireinject
// +build wireinject

package module

import (
	"common/pkg/client/rpc"
	"log/slog"

	"github.com/google/wire"
)

func New(*Config, *slog.Logger, *rpc.UserClient, *rpc.ContentClient) (*Module, func(), error) {
	panic(wire.Build(ProviderSet))
}
