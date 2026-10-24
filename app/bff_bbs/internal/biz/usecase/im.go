package usecase

import (
	"context"

	"common/pkg/client/rpc"
	"common/pkg/constant"
	imv1 "common/proto/gen/im/v1"
	realtimehubv1 "common/proto/gen/realtime_hub/v1"

	"bff_bbs/internal/biz/model"
	"bff_bbs/internal/biz/repo"
)

type IMUsecase struct {
	privateChat       imv1.IMPrivateChatServiceClient
	globalChat        imv1.IMGlobalChatServiceClient
	globalChatBarrage imv1.IMGlobalChatBarrageServiceClient
	accounts          repo.AccountClient
	assets            repo.AssetClient
	hub               realtimehubv1.RealtimeHubServiceClient
}

func NewIMUsecase(client *rpc.IMClient, accounts repo.AccountClient, assets repo.AssetClient, hub *rpc.RealtimeHubClient) *IMUsecase {
	return &IMUsecase{
		privateChat:       client.PrivateChat,
		globalChat:        client.GlobalChat,
		globalChatBarrage: client.GlobalChatBarrage,
		accounts:          accounts,
		assets:            assets,
		hub:               hub.Hub,
	}
}

func (u *IMUsecase) MapChatProfiles(ctx context.Context, userIDs []int64) (map[int64]*model.AccountProfile, error) {
	profiles, err := u.accounts.MapProfileAccounts(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	for _, profile := range profiles {
		if profile == nil {
			continue
		}
		avatarURL := "/v1/user/account/avatar?name=" + profile.Name
		if profile.AvatarAssetID != nil && *profile.AvatarAssetID > 0 {
			asset, assetErr := u.assets.Get(ctx, *profile.AvatarAssetID)
			if assetErr != nil {
				return nil, assetErr
			}
			if asset != nil && asset.URL != "" {
				avatarURL = asset.URL
			}
		}
		profile.AvatarURL = &avatarURL
	}
	return profiles, nil
}

func (u *IMUsecase) ListGlobalChatOnlineMembers(ctx context.Context) ([]*model.AccountProfile, error) {
	response, err := u.hub.ListOnlineUsers(ctx, &realtimehubv1.ListRealtimeOnlineUsers_Req{Realm: constant.RealtimeRealmBBS})
	if err != nil {
		return nil, err
	}
	profiles, err := u.MapChatProfiles(ctx, response.GetUserIds())
	if err != nil {
		return nil, err
	}
	rows := make([]*model.AccountProfile, 0, len(response.GetUserIds()))
	for _, userID := range response.GetUserIds() {
		if profile := profiles[userID]; profile != nil {
			rows = append(rows, profile)
		}
	}
	return rows, nil
}

func (u *IMUsecase) QueryChatOnline(ctx context.Context, userIDs []int64) (map[int64]bool, error) {
	response, err := u.hub.QueryOnline(ctx, &realtimehubv1.QueryRealtimeOnline_Req{Realm: constant.RealtimeRealmBBS, UserIds: userIDs})
	if err != nil {
		return nil, err
	}
	return response.GetOnlineStatuses(), nil
}

func (u *IMUsecase) ListPrivateChatSessions(ctx context.Context, req *imv1.ListPrivateChatSessions_Req) (*imv1.ListPrivateChatSessions_Resp, error) {
	return u.privateChat.ListSessions(ctx, req)
}

func (u *IMUsecase) ListPrivateChatMessages(ctx context.Context, req *imv1.ListPrivateChatMessages_Req) (*imv1.ListPrivateChatMessages_Resp, error) {
	return u.privateChat.ListMessages(ctx, req)
}

func (u *IMUsecase) CountPrivateChatUnread(ctx context.Context, userID int64) (uint32, error) {
	response, err := u.privateChat.CountUnread(ctx, &imv1.CountUnreadPrivateChatMessages_Req{UserId: userID})
	if err != nil {
		return 0, err
	}
	return response.GetCount(), nil
}

func (u *IMUsecase) SendPrivateChatMessage(ctx context.Context, req *imv1.SendPrivateChatMessage_Req) (*imv1.SendPrivateChatMessage_Resp, error) {
	return u.privateChat.Send(ctx, req)
}

func (u *IMUsecase) RevokePrivateChatMessage(ctx context.Context, req *imv1.RevokePrivateChatMessage_Req) (*imv1.RevokePrivateChatMessage_Resp, error) {
	return u.privateChat.Revoke(ctx, req)
}

func (u *IMUsecase) MarkPrivateChatSessionRead(ctx context.Context, req *imv1.MarkPrivateChatSessionRead_Req) error {
	_, err := u.privateChat.MarkRead(ctx, req)
	return err
}

func (u *IMUsecase) SetPrivateChatSessionMuted(ctx context.Context, req *imv1.SetPrivateChatSessionMuted_Req) error {
	_, err := u.privateChat.SetMuted(ctx, req)
	return err
}

func (u *IMUsecase) SetPrivateChatSessionPinned(ctx context.Context, req *imv1.SetPrivateChatSessionPinned_Req) error {
	_, err := u.privateChat.SetPinned(ctx, req)
	return err
}

func (u *IMUsecase) ListGlobalChatMessages(ctx context.Context, req *imv1.ListGlobalChatMessages_Req) (*imv1.ListGlobalChatMessages_Resp, error) {
	return u.globalChat.List(ctx, req)
}

func (u *IMUsecase) GetGlobalChatTopic(ctx context.Context) (*imv1.GetGlobalChatTopic_Resp, error) {
	return u.globalChat.GetTopic(ctx, &imv1.GetGlobalChatTopic_Req{})
}

func (u *IMUsecase) SetGlobalChatTopic(ctx context.Context, req *imv1.SetGlobalChatTopic_Req) (*imv1.SetGlobalChatTopic_Resp, error) {
	return u.globalChat.SetTopic(ctx, req)
}

func (u *IMUsecase) SendGlobalChatMessage(ctx context.Context, req *imv1.SendGlobalChatMessage_Req) (*imv1.SendGlobalChatMessage_Resp, error) {
	return u.globalChat.Send(ctx, req)
}

func (u *IMUsecase) SendGlobalChatBarrage(ctx context.Context, req *imv1.SendGlobalChatBarrage_Req) (*imv1.SendGlobalChatBarrage_Resp, error) {
	return u.globalChatBarrage.Send(ctx, req)
}

func (u *IMUsecase) ListGlobalChatBarrages(ctx context.Context, req *imv1.ListGlobalChatBarrages_Req) (*imv1.ListGlobalChatBarrages_Resp, error) {
	return u.globalChatBarrage.List(ctx, req)
}

func (u *IMUsecase) RevokeGlobalChatMessage(ctx context.Context, req *imv1.RevokeGlobalChatMessage_Req) (*imv1.RevokeGlobalChatMessage_Resp, error) {
	return u.globalChat.Revoke(ctx, req)
}
