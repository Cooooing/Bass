//go:build wireinject
// +build wireinject

package module

import (
	"log/slog"

	"github.com/google/wire"
)

func New(*Config, *slog.Logger) (*Module, func(), error) {
	panic(wire.Build(ProviderSet))
}
