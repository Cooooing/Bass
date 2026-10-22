package usecase

import (
	"context"

	"common/pkg/client/rpc"
	"common/pkg/constant"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"
)

type RealtimeUsecase struct {
	hub realtimehubv1.RealtimeHubServiceClient
}

func NewRealtimeUsecase(hub *rpc.RealtimeHubClient) *RealtimeUsecase {
	return &RealtimeUsecase{hub: hub.Hub}
}

func (u *RealtimeUsecase) CreateTicket(ctx context.Context, userID int64) (string, int64, error) {
	resp, err := u.hub.IssueTicket(ctx, &realtimehubv1.IssueRealtimeTicket_Req{UserId: userID, Realm: constant.RealtimeRealmBBS})
	if err != nil {
		return "", 0, err
	}
	return resp.GetTicket(), resp.GetExpiresInSeconds(), nil
}
