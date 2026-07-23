package rpc

import (
	commonv1 "common/proto/gen/common/v1"
	gametownv1 "common/proto/gen/game_town/v1"

	"google.golang.org/grpc"
)

type GameTownClient struct {
	System      commonv1.CommonSystemServiceClient
	AgentConfig gametownv1.GameTownAgentConfigServiceClient
	Player      gametownv1.GameTownPlayerServiceClient
	World       gametownv1.GameTownWorldServiceClient
	WorldMember gametownv1.GameTownWorldMemberServiceClient
	Npc         gametownv1.GameTownNpcServiceClient
	Location    gametownv1.GameTownLocationServiceClient
	Faction     gametownv1.GameTownFactionServiceClient
	Event       gametownv1.GameTownEventServiceClient
}

func NewGameTownClient(conn *grpc.ClientConn) *GameTownClient {
	return &GameTownClient{
		System:      commonv1.NewCommonSystemServiceClient(conn),
		AgentConfig: gametownv1.NewGameTownAgentConfigServiceClient(conn),
		Player:      gametownv1.NewGameTownPlayerServiceClient(conn),
		World:       gametownv1.NewGameTownWorldServiceClient(conn),
		WorldMember: gametownv1.NewGameTownWorldMemberServiceClient(conn),
		Npc:         gametownv1.NewGameTownNpcServiceClient(conn),
		Location:    gametownv1.NewGameTownLocationServiceClient(conn),
		Faction:     gametownv1.NewGameTownFactionServiceClient(conn),
		Event:       gametownv1.NewGameTownEventServiceClient(conn),
	}
}
