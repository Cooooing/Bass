//go:build wireinject

package module

import (
	commonmodule "common/pkg/module"
	"github.com/google/wire"
	"log/slog"
)

func wireModule(*Config, *slog.Logger, *commonmodule.Infrastructure) (*Module, error) {
	panic(wire.Build(ProviderSet))
}
