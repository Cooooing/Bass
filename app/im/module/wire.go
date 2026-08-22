//go:build wireinject
// +build wireinject

package module

import (
	"log/slog"

	"github.com/google/wire"
)

func wireModule(*Config, *slog.Logger) (*Module, func(), error) {
	panic(wire.Build(ProviderSet))
}
