package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	bbsuserv1enum "common/proto/gen/bbs/v1/user/enum"
	"context"
)

type AccountUsecase struct {
	accountClient repo.AccountClient
}

func NewAccountUsecase(
	accountClient repo.AccountClient,
) *AccountUsecase {
	return &AccountUsecase{
		accountClient: accountClient,
	}
}

func (u *AccountUsecase) GetCurrentAccount(ctx context.Context, userID int64) (*bbsuserv1.GetCurrentAccount_Resp_Account, error) {
	reply, err := u.accountClient.GetCurrentAccount(ctx, userID)
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.GetCurrentAccount_Resp_Account
	if reply != nil {
		account = &bbsuserv1.GetCurrentAccount_Resp_Account{}
		if profile := reply.Profile; profile != nil {
			account.Profile = &bbsuserv1.GetCurrentAccount_Resp_AccountProfile{
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
				CreatedAt:     profile.CreatedAt,
				UpdatedAt:     profile.UpdatedAt,
			}
		}
		if contact := reply.Contact; contact != nil {
			account.Contact = &bbsuserv1.GetCurrentAccount_Resp_AccountContact{
				UserId: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return account, nil
}

func (u *AccountUsecase) GetProfileAccount(ctx context.Context, userID int64) (*bbsuserv1.GetProfileAccount_Resp_AccountProfile, error) {
	reply, err := u.accountClient.GetProfileAccount(ctx, userID)
	if err != nil {
		return nil, err
	}
	var profile *bbsuserv1.GetProfileAccount_Resp_AccountProfile
	if row := reply; row != nil {
		profile = &bbsuserv1.GetProfileAccount_Resp_AccountProfile{
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
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		}
	}
	return profile, nil
}

type UpdateProfileAccountReq struct {
	UserID       int64
	AvatarURL    *string
	Nickname     *string
	URL          *string
	Introduction *string
	Mbti         *bbsuserv1enum.MBTI
}

func (u *AccountUsecase) UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*bbsuserv1.UpdateProfileAccount_Resp_AccountProfile, error) {
	var mbti *int32
	if req.Mbti != nil {
		value := int32(*req.Mbti)
		mbti = &value
	}
	reply, err := u.accountClient.UpdateProfileAccount(ctx, &repo.UpdateProfileAccountReq{
		UserID:       req.UserID,
		AvatarURL:    req.AvatarURL,
		Nickname:     req.Nickname,
		URL:          req.URL,
		Introduction: req.Introduction,
		MBTI:         mbti,
	})
	if err != nil {
		return nil, err
	}
	var profile *bbsuserv1.UpdateProfileAccount_Resp_AccountProfile
	if row := reply; row != nil {
		profile = &bbsuserv1.UpdateProfileAccount_Resp_AccountProfile{
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
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		}
	}
	return profile, nil
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
