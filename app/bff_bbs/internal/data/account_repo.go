package data

import (
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"time"

	"bff_bbs/internal/biz/model"

	"bff_bbs/internal/biz/repo"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ repo.AccountClient = (*AccountClient)(nil)

type AccountClient struct {
	userClient *rpc.UserClient
}

func NewAccountClient(
	userClient *rpc.UserClient,
) repo.AccountClient {
	return &AccountClient{
		userClient: userClient,
	}
}
func (r *AccountClient) GetCurrentAccount(ctx context.Context, userID int64) (*model.Account, error) {
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Req{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var out *model.Account
	if account != nil {
		out = &model.Account{}
		if basic := account.GetBasic(); basic != nil {
			out.Profile = &model.AccountProfile{
				ID:                basic.GetId(),
				Name:              basic.GetName(),
				Nickname:          basic.Nickname,
				URL:               basic.Url,
				AvatarAssetID:     basic.AvatarAssetId,
				BackgroundAssetID: basic.BackgroundAssetId,
				Introduction:      basic.Introduction,
				Status:            int32(basic.GetStatus()),
				MBTI:              basic.Mbti,
				FollowCount:       basic.FollowCount,
				FollowerCount:     basic.FollowerCount,
				CreatedAt:         new(basic.GetCreatedAt().AsTime()),
				UpdatedAt:         new(basic.GetUpdatedAt().AsTime()),
			}
		}
		if contact := account.GetContact(); contact != nil {
			out.Contact = &model.AccountContact{
				UserID: contact.GetUserId(),
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return out, nil
}

func (r *AccountClient) GetProfile(ctx context.Context, name string) (*model.AccountProfile, *time.Time, error) {
	reply, err := r.userClient.Account.Get(ctx, &userv1.GetAccount_Req{Name: &name})
	if err != nil {
		return nil, nil, err
	}
	return r.accountProfile(reply.GetAccount().GetBasic()), r.timestampValue(reply.LastSuccessLoginAt), nil
}

func (r *AccountClient) MapProfileAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.AccountProfile, error) {
	if len(userIDs) == 0 {
		return map[int64]*model.AccountProfile{}, nil
	}
	reply, err := r.userClient.Account.Map(ctx, &userv1.MapAccounts_Req{Query: &userv1.MapAccounts_Req_AccountQuery{UserIds: userIDs}})
	if err != nil {
		return nil, err
	}
	profiles := make(map[int64]*model.AccountProfile, len(reply.GetAccounts()))
	for id, account := range reply.GetAccounts() {
		profiles[id] = r.accountProfile(account.GetBasic())
	}
	return profiles, nil
}

func (r *AccountClient) accountProfile(account *userv1.AccountBasic) *model.AccountProfile {
	if account == nil {
		return nil
	}
	return &model.AccountProfile{ID: account.GetId(), Name: account.GetName(), Nickname: account.Nickname, URL: account.Url, AvatarAssetID: account.AvatarAssetId, BackgroundAssetID: account.BackgroundAssetId, Introduction: account.Introduction, Status: int32(account.GetStatus()), MBTI: account.Mbti, FollowCount: account.FollowCount, FollowerCount: account.FollowerCount, CreatedAt: r.timestampValue(account.GetCreatedAt()), UpdatedAt: r.timestampValue(account.GetUpdatedAt())}
}

func (*AccountClient) timestampValue(value *timestamppb.Timestamp) *time.Time {
	if value == nil {
		return nil
	}
	result := value.AsTime()
	return &result
}

func (r *AccountClient) UpdateProfileAccount(ctx context.Context, req *repo.UpdateProfileAccountReq) (*model.AccountProfile, error) {
	updateReq := &userv1.UpdateProfileAccount_Req{
		UserId:            req.UserID,
		AvatarAssetId:     req.AvatarAssetID,
		BackgroundAssetId: req.BackgroundAssetID,
		Nickname:          req.Nickname,
		Url:               req.URL,
		Introduction:      req.Introduction,
	}
	updateReq.Mbti = req.MBTI
	reply, err := r.userClient.Account.UpdateProfile(ctx, updateReq)
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var profile *model.AccountProfile
	if account != nil {
		profile = &model.AccountProfile{
			ID:                account.GetId(),
			Name:              account.GetName(),
			Nickname:          account.Nickname,
			URL:               account.Url,
			AvatarAssetID:     account.AvatarAssetId,
			BackgroundAssetID: account.BackgroundAssetId,
			Introduction:      account.Introduction,
			Status:            int32(account.GetStatus()),
			MBTI:              account.Mbti,
			FollowCount:       account.FollowCount,
			FollowerCount:     account.FollowerCount,
			CreatedAt:         new(account.GetCreatedAt().AsTime()),
			UpdatedAt:         new(account.GetUpdatedAt().AsTime()),
		}
	}
	return profile, nil
}

func (r *AccountClient) UpdatePasswordAccount(ctx context.Context, req *repo.UpdatePasswordAccountReq) error {
	_, err := r.userClient.Account.UpdatePassword(ctx, &userv1.UpdatePasswordAccount_Req{
		UserId:      req.UserID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
	return err
}

func (r *AccountClient) UpdateEmailAccount(ctx context.Context, req *repo.UpdateEmailAccountReq) error {
	_, err := r.userClient.Account.UpdateEmail(ctx, &userv1.UpdateEmailAccount_Req{
		UserId: req.UserID,
		Email:  req.Email,
		Code:   req.Code,
	})
	return err
}

func (r *AccountClient) UpdatePhoneAccount(ctx context.Context, req *repo.UpdatePhoneAccountReq) error {
	_, err := r.userClient.Account.UpdatePhone(ctx, &userv1.UpdatePhoneAccount_Req{
		UserId: req.UserID,
		Phone:  req.Phone,
		Code:   req.Code,
	})
	return err
}
func (r *AccountClient) AvatarAccount(ctx context.Context, name string) (*repo.AvatarAccountResp, error) {
	reply, err := r.userClient.Account.Avatar(ctx, &userv1.AvatarAccount_Req{
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	return &repo.AvatarAccountResp{
		Data:        reply.GetData(),
		ContentType: reply.GetContentType(),
	}, nil
}
