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
	NewWebSocketService,
	NewWebSocketTicketService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	authService *AuthService,
	characterService *CharacterService,
	webSocketService *WebSocketService,
	webSocketTicketService *WebSocketTicketService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		authService,
		characterService,
		webSocketService,
		webSocketTicketService,
	}
}
