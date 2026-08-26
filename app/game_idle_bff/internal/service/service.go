package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewAuthService,
	NewCharacterService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	authService *AuthService,
	characterService *CharacterService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		authService,
		characterService,
	}
}
