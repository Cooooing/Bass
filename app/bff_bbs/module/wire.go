//go:build wireinject
// +build wireinject

package module

import (
	commonmodule "common/pkg/module"
	"log/slog"

	"github.com/google/wire"
)

func wireModule(*Config, *slog.Logger, *commonmodule.Clients, *commonmodule.Infrastructure) (*Module, error) {
	panic(wire.Build(ProviderSet))
}
