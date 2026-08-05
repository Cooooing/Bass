package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	bbsuserv1enum "common/proto/gen/bbs/v1/user/enum"
	economyv1enum "common/proto/gen/economy/v1/enum"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountUsecase struct {
	accountClient repo.AccountClient
	assetClient   repo.AssetClient
	economyClient repo.EconomyClient
}

func NewAccountUsecase(
	accountClient repo.AccountClient,
	assetClient repo.AssetClient,
	economyClient repo.EconomyClient,
) *AccountUsecase {
	return &AccountUsecase{
		accountClient: accountClient,
		assetClient:   assetClient,
		economyClient: economyClient,
	}
}
func (u *AccountUsecase) GetCurrentAccount(ctx context.Context, userID int64) (*bbsuserv1.GetCurrentAccount_Resp_Account, error) {
	reply, err := u.accountClient.GetCurrentAccount(ctx, userID)
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.GetCurrentAccount_Resp_Account
	if reply != nil {
		if reply.Profile != nil {
			avatarURL := "/v1/user/account/avatar?name=" + reply.Profile.Name
			if reply.Profile.AvatarAssetID != nil && *reply.Profile.AvatarAssetID > 0 && u.assetClient != nil {
				asset, err := u.assetClient.Get(ctx, *reply.Profile.AvatarAssetID)
				if err != nil {
					return nil, err
				}
				if asset != nil && asset.URL != "" {
					avatarURL = asset.URL
				}
			}
			reply.Profile.AvatarURL = &avatarURL
		}
		account = &bbsuserv1.GetCurrentAccount_Resp_Account{}
		if profile := reply.Profile; profile != nil {
			account.Profile = &bbsuserv1.AccountProfile{
				Id:            profile.ID,
				Name:          profile.Name,
				Nickname:      profile.Nickname,
				Url:           profile.URL,
				AvatarUrl:     profile.AvatarURL,
				Introduction:  profile.Introduction,
				Status:        bbsuserv1enum.AccountStatus(profile.Status),
				Mbti:          bbsuserv1enum.MBTI(profile.MBTI),
				FollowCount:   profile.FollowCount,
				FollowerCount: profile.FollowerCount,
			}
			if profile.CreatedAt != nil {
				account.Profile.CreatedAt = timestamppb.New(*profile.CreatedAt)
			}
			if profile.UpdatedAt != nil {
				account.Profile.UpdatedAt = timestamppb.New(*profile.UpdatedAt)
			}
		}
		if contact := reply.Contact; contact != nil {
			account.Contact = &bbsuserv1.AccountContact{
				UserId: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return account, nil
}

func (u *AccountUsecase) GetProfileAccount(ctx context.Context, userID int64) (*bbsuserv1.AccountProfile, error) {
	reply, err := u.accountClient.GetProfileAccount(ctx, userID)
	if err != nil {
		return nil, err
	}
	var profile *bbsuserv1.AccountProfile
	if row := reply; row != nil {
		avatarURL := "/v1/user/account/avatar?name=" + row.Name
		if row.AvatarAssetID != nil && *row.AvatarAssetID > 0 && u.assetClient != nil {
			asset, err := u.assetClient.Get(ctx, *row.AvatarAssetID)
			if err != nil {
				return nil, err
			}
			if asset != nil && asset.URL != "" {
				avatarURL = asset.URL
			}
		}
		row.AvatarURL = &avatarURL
		profile = &bbsuserv1.AccountProfile{
			Id:            row.ID,
			Name:          row.Name,
			Nickname:      row.Nickname,
			Url:           row.URL,
			AvatarUrl:     row.AvatarURL,
			Introduction:  row.Introduction,
			Status:        bbsuserv1enum.AccountStatus(row.Status),
			Mbti:          bbsuserv1enum.MBTI(row.MBTI),
			FollowCount:   row.FollowCount,
			FollowerCount: row.FollowerCount,
		}
		if row.CreatedAt != nil {
			profile.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			profile.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
	}
	return profile, nil
}

type UpdateProfileAccountReq struct {
	UserID        int64
	AvatarAssetID *int64
	Nickname      *string
	URL           *string
	Introduction  *string
	Mbti          *bbsuserv1enum.MBTI
}

func (u *AccountUsecase) UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*bbsuserv1.AccountProfile, error) {
	var mbti *int32
	if req.Mbti != nil {
		mbti = new(int32(*req.Mbti))
	}
	reply, err := u.accountClient.UpdateProfileAccount(ctx, &repo.UpdateProfileAccountReq{
		UserID:        req.UserID,
		AvatarAssetID: req.AvatarAssetID,
		Nickname:      req.Nickname,
		URL:           req.URL,
		Introduction:  req.Introduction,
		MBTI:          mbti,
	})
	if err != nil {
		return nil, err
	}
	var profile *bbsuserv1.AccountProfile
	if row := reply; row != nil {
		avatarURL := "/v1/user/account/avatar?name=" + row.Name
		if row.AvatarAssetID != nil && *row.AvatarAssetID > 0 && u.assetClient != nil {
			asset, err := u.assetClient.Get(ctx, *row.AvatarAssetID)
			if err != nil {
				return nil, err
			}
			if asset != nil && asset.URL != "" {
				avatarURL = asset.URL
			}
		}
		row.AvatarURL = &avatarURL
		profile = &bbsuserv1.AccountProfile{
			Id:            row.ID,
			Name:          row.Name,
			Nickname:      row.Nickname,
			Url:           row.URL,
			AvatarUrl:     row.AvatarURL,
			Introduction:  row.Introduction,
			Status:        bbsuserv1enum.AccountStatus(row.Status),
			Mbti:          bbsuserv1enum.MBTI(row.MBTI),
			FollowCount:   row.FollowCount,
			FollowerCount: row.FollowerCount,
		}
		if row.CreatedAt != nil {
			profile.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			profile.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
	}
	return profile, nil
}

type UpdatePasswordAccountReq struct {
	UserID      int64
	OldPassword string
	NewPassword string
}

func (u *AccountUsecase) UpdatePasswordAccount(ctx context.Context, req *UpdatePasswordAccountReq) error {
	return u.accountClient.UpdatePasswordAccount(ctx, &repo.UpdatePasswordAccountReq{
		UserID:      req.UserID,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
	})
}

type UpdateEmailAccountReq struct {
	UserID int64
	Email  string
	Code   string
}

func (u *AccountUsecase) UpdateEmailAccount(ctx context.Context, req *UpdateEmailAccountReq) error {
	return u.accountClient.UpdateEmailAccount(ctx, &repo.UpdateEmailAccountReq{
		UserID: req.UserID,
		Email:  req.Email,
		Code:   req.Code,
	})
}

type UpdatePhoneAccountReq struct {
	UserID int64
	Phone  string
	Code   string
}

func (u *AccountUsecase) UpdatePhoneAccount(ctx context.Context, req *UpdatePhoneAccountReq) error {
	return u.accountClient.UpdatePhoneAccount(ctx, &repo.UpdatePhoneAccountReq{
		UserID: req.UserID,
		Phone:  req.Phone,
		Code:   req.Code,
	})
}

type AvatarAccountResp struct {
	Data        []byte
	ContentType string
}

func (u *AccountUsecase) AvatarAccount(ctx context.Context, name string) (*AvatarAccountResp, error) {
	reply, err := u.accountClient.AvatarAccount(ctx, name)
	if err != nil {
		return nil, err
	}
	return &AvatarAccountResp{
		Data:        reply.Data,
		ContentType: reply.ContentType,
	}, nil
}

type AccountEconomyResp struct {
	Balance      int64
	TotalIncome  int64
	TotalExpense int64
}

func (u *AccountUsecase) GetEconomyAccount(ctx context.Context, userID int64) (*AccountEconomyResp, error) {
	account, err := u.economyClient.GetAccount(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &AccountEconomyResp{Balance: account.Balance, TotalIncome: account.TotalIncome, TotalExpense: account.TotalExpense}, nil
}

type ListAccountEconomyRecordsReq struct {
	UserID     int64
	Page       *repo.PageReq
	Direction  *economyv1enum.EconomyRecordDirection
	RecordType *economyv1enum.EconomyRecordType
}

type ListAccountEconomyRecordsResp struct {
	Rows []*repo.EconomyRecord
	Page *repo.PageResp
}

func (u *AccountUsecase) ListEconomyRecords(ctx context.Context, req *ListAccountEconomyRecordsReq) (*ListAccountEconomyRecordsResp, error) {
	resp, err := u.economyClient.ListRecords(ctx, &repo.ListEconomyRecordsReq{UserID: req.UserID, Page: req.Page, Direction: req.Direction, RecordType: req.RecordType})
	if err != nil {
		return nil, err
	}
	return &ListAccountEconomyRecordsResp{Rows: resp.Rows, Page: resp.Page}, nil
}
