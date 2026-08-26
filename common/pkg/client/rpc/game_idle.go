package rpc

import (
	"common/pkg/client/localrpc"
	commonv1 "common/proto/gen/common/v1"
	gameidlev1 "common/proto/gen/game_idle/v1"

	"google.golang.org/grpc"
)

type GameIdleClient struct {
	System      commonv1.CommonSystemServiceClient
	Character   gameidlev1.GameIdleCharacterServiceClient
	Backpack    gameidlev1.GameIdleBackpackServiceClient
	ActionQueue gameidlev1.GameIdleActionQueueServiceClient
	Chat        gameidlev1.GameIdleChatServiceClient
}

func NewGameIdleClient(conn grpc.ClientConnInterface) *GameIdleClient {
	return &GameIdleClient{
		System:      commonv1.NewCommonSystemServiceClient(conn),
		Character:   gameidlev1.NewGameIdleCharacterServiceClient(conn),
		Backpack:    gameidlev1.NewGameIdleBackpackServiceClient(conn),
		ActionQueue: gameidlev1.NewGameIdleActionQueueServiceClient(conn),
		Chat:        gameidlev1.NewGameIdleChatServiceClient(conn),
	}
}

func MountGameIdleServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&commonv1.CommonSystemService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.GameIdleCharacterService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.GameIdleBackpackService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.GameIdleActionQueueService_ServiceDesc, service)
		conn.RegisterMatching(&gameidlev1.GameIdleChatService_ServiceDesc, service)
	}
}
