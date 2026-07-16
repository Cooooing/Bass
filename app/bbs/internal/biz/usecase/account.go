package usecase

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/proto/gen/bbs/v1/user"
	"context"
)

type AccountUsecase struct {
	accountClient repo.AccountClient
}

func NewAccountUsecase(accountClient repo.AccountClient) *AccountUsecase {
	return &AccountUsecase{accountClient: accountClient}
}

type GetCurrentAccountReq struct {
	UserID int64
}

type GetCurrentAccountResponse struct {
	Account *bbsuserv1.GetCurrentAccount_Response_Account
}

func (u *AccountUsecase) GetCurrentAccount(ctx context.Context, req *GetCurrentAccountReq) (*GetCurrentAccountResponse, error) {
	reply, err := u.accountClient.GetCurrentAccount(ctx, &repo.GetCurrentAccountReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var account *bbsuserv1.GetCurrentAccount_Response_Account
	if reply.Account != nil {
		account = &bbsuserv1.GetCurrentAccount_Response_Account{}
		if profile := reply.Account.Profile; profile != nil {
			account.Profile = &bbsuserv1.GetCurrentAccount_Response_AccountProfile{
				Id:            profile.ID,
				Name:          profile.Name,
				Nickname:      profile.Nickname,
				Url:           profile.URL,
				AvatarUrl:     profile.AvatarURL,
				Introduction:  profile.Introduction,
				Status:        bbsuserv1.AccountStatus(profile.Status),
				Mbti:          bbsuserv1.MBTI(profile.MBTI),
				FollowCount:   profile.FollowCount,
				FollowerCount: profile.FollowerCount,
				CreatedAt:     profile.CreatedAt,
				UpdatedAt:     profile.UpdatedAt,
			}
		}
		if contact := reply.Account.Contact; contact != nil {
			account.Contact = &bbsuserv1.GetCurrentAccount_Response_AccountContact{
				UserId: contact.UserID,
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &GetCurrentAccountResponse{Account: account}, nil
}

type GetProfileAccountReq struct {
	UserID int64
}

type GetProfileAccountResponse struct {
	Profile *bbsuserv1.GetProfileAccount_Response_AccountProfile
}

func (u *AccountUsecase) GetProfileAccount(ctx context.Context, req *GetProfileAccountReq) (*GetProfileAccountResponse, error) {
	reply, err := u.accountClient.GetProfileAccount(ctx, &repo.GetProfileAccountReq{UserID: req.UserID})
	if err != nil {
		return nil, err
	}
	var profile *bbsuserv1.GetProfileAccount_Response_AccountProfile
	if row := reply.Profile; row != nil {
		profile = &bbsuserv1.GetProfileAccount_Response_AccountProfile{
			Id:            row.ID,
			Name:          row.Name,
			Nickname:      row.Nickname,
			Url:           row.URL,
			AvatarUrl:     row.AvatarURL,
			Introduction:  row.Introduction,
			Status:        bbsuserv1.AccountStatus(row.Status),
			Mbti:          bbsuserv1.MBTI(row.MBTI),
			FollowCount:   row.FollowCount,
			FollowerCount: row.FollowerCount,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		}
	}
	return &GetProfileAccountResponse{Profile: profile}, nil
}

type UpdateProfileAccountReq struct {
	UserID       int64
	AvatarURL    *string
	Nickname     *string
	URL          *string
	Introduction *string
	Mbti         *bbsuserv1.MBTI
}

type UpdateProfileAccountResponse struct {
	Profile *bbsuserv1.UpdateProfileAccount_Response_AccountProfile
}

func (u *AccountUsecase) UpdateProfileAccount(ctx context.Context, req *UpdateProfileAccountReq) (*UpdateProfileAccountResponse, error) {
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
	var profile *bbsuserv1.UpdateProfileAccount_Response_AccountProfile
	if row := reply.Profile; row != nil {
		profile = &bbsuserv1.UpdateProfileAccount_Response_AccountProfile{
			Id:            row.ID,
			Name:          row.Name,
			Nickname:      row.Nickname,
			Url:           row.URL,
			AvatarUrl:     row.AvatarURL,
			Introduction:  row.Introduction,
			Status:        bbsuserv1.AccountStatus(row.Status),
			Mbti:          bbsuserv1.MBTI(row.MBTI),
			FollowCount:   row.FollowCount,
			FollowerCount: row.FollowerCount,
			CreatedAt:     row.CreatedAt,
			UpdatedAt:     row.UpdatedAt,
		}
	}
	return &UpdateProfileAccountResponse{Profile: profile}, nil
}

type AvatarAccountReq struct {
	Name string
}

type AvatarAccountResponse struct {
	Data        []byte
	ContentType string
}

func (u *AccountUsecase) AvatarAccount(ctx context.Context, req *AvatarAccountReq) (*AvatarAccountResponse, error) {
	reply, err := u.accountClient.AvatarAccount(ctx, &repo.AvatarAccountReq{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &AvatarAccountResponse{Data: reply.Data, ContentType: reply.ContentType}, nil
}
