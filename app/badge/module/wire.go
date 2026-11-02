//go:build wireinject
// +build wireinject

package module

import (
	"log/slog"

	commonmodule "common/pkg/module"

	"github.com/google/wire"
)

func wireModule(*Config, *slog.Logger, *commonmodule.Clients) (*Module, func(), error) {
	panic(wire.Build(ProviderSet))
}
