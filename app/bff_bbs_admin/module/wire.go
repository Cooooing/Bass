//go:build wireinject
// +build wireinject

package module

import (
	"github.com/google/wire"
)

func wireModule(*Config) (*Module, error) { panic(wire.Build(ProviderSet)) }
