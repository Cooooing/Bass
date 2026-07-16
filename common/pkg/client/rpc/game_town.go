package rpc

import (
	gametownv1 "common/proto/gen/game_town/v1"

	"google.golang.org/grpc"
)

type GameTownClient struct {
	AgentConfig gametownv1.GameTownAgentConfigServiceClient
	World       gametownv1.GameTownWorldServiceClient
	Npc         gametownv1.GameTownNpcServiceClient
	Command     gametownv1.GameTownCommandServiceClient
	Event       gametownv1.GameTownEventServiceClient
	Memory      gametownv1.GameTownMemoryServiceClient
	Session     gametownv1.GameTownSessionServiceClient
	Player      gametownv1.GameTownPlayerServiceClient
}

func NewGameTownClient(conn *grpc.ClientConn) *GameTownClient {
	return &GameTownClient{
		AgentConfig: gametownv1.NewGameTownAgentConfigServiceClient(conn),
		World:       gametownv1.NewGameTownWorldServiceClient(conn),
		Npc:         gametownv1.NewGameTownNpcServiceClient(conn),
		Command:     gametownv1.NewGameTownCommandServiceClient(conn),
		Event:       gametownv1.NewGameTownEventServiceClient(conn),
		Memory:      gametownv1.NewGameTownMemoryServiceClient(conn),
		Session:     gametownv1.NewGameTownSessionServiceClient(conn),
		Player:      gametownv1.NewGameTownPlayerServiceClient(conn),
	}
}
