//go:build wireinject
// +build wireinject

package module

import (
	commonmodule "common/pkg/module"

	"github.com/google/wire"
)

func wireModule(*Config, *commonmodule.Clients) (*Module, error) {
	panic(wire.Build(ProviderSet))
}
