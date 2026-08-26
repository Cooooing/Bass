//go:build wireinject
// +build wireinject

package module

import (
	commonmodule "common/pkg/module"
	"log/slog"

	"github.com/google/wire"
)

func wireModule(
	config *Config,
	logger *slog.Logger,
	infrastructure *commonmodule.Infrastructure,
) (*Module, func(), error) {
	panic(wire.Build(ProviderSet))
}
