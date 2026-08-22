package rpc

import (
	"common/pkg/client/localrpc"
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

func NewGameTownClient(
	conn grpc.ClientConnInterface,
) *GameTownClient {
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

func NewLocalGameTownClient[T any](services []T) *GameTownClient {
	conn := localrpc.NewConn()
	MountGameTownServices(conn, services)
	return NewGameTownClient(conn)
}

func MountGameTownServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&commonv1.CommonSystemService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownAgentConfigService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownPlayerService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownWorldService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownWorldMemberService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownNpcService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownLocationService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownFactionService_ServiceDesc, service)
		conn.RegisterMatching(&gametownv1.GameTownEventService_ServiceDesc, service)
	}
}
