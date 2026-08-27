package rpc

import (
	"common/pkg/client/localrpc"
	commonv1 "common/proto/gen/common/v1"
	gameidlev1 "common/proto/gen/game_idle/v1"

	"google.golang.org/grpc"
)

type GameIdleClient struct {
	System      commonv1.CommonSystemServiceClient
	Character   gameidlev1.CharacterServiceClient
	Ability     gameidlev1.CharacterAbilityServiceClient
	Backpack    gameidlev1.BackpackServiceClient
	ActionQueue gameidlev1.ActionQueueServiceClient
	Chat        gameidlev1.ChatServiceClient
}

func NewGameIdleClient(conn grpc.ClientConnInterface) *GameIdleClient {
	return &GameIdleClient{
		System:      commonv1.NewCommonSystemServiceClient(conn),
		Character:   gameidlev1.NewCharacterServiceClient(conn),
		Ability:     gameidlev1.NewCharacterAbilityServiceClient(conn),
		Backpack:    gameidlev1.NewBackpackServiceClient(conn),
		ActionQueue: gameidlev1.NewActionQueueServiceClient(conn),
		Chat:        gameidlev1.NewChatServiceClient(conn),
	}
}

func MountGameIdleServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&commonv1.CommonSystemService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.CharacterService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.CharacterAbilityService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.BackpackService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.ActionQueueService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.ChatService_ServiceDesc, service)
	}
}
