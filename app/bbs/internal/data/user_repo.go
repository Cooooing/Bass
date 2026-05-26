package data

import (
	"bbs/internal/biz/repo"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	commonv1 "common/api/gen/common"
	userv1 "common/api/gen/user/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	userClient *rpc.UserClient
}

func NewUserRepo(userClient *rpc.UserClient) repo.UserRepo {
	return &UserRepo{userClient: userClient}
}

func (r *UserRepo) StartEmailRegistration(ctx context.Context, req *bbsuserv1.StartEmailRegistration_Request) (*bbsuserv1.StartEmailRegistration_Reply, error) {
	reply, err := r.userClient.Auth.StartEmailRegistration(ctx, &userv1.StartEmailRegistration_Request{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.StartEmailRegistration_Reply{CodeToken: reply.GetCodeToken()}, nil
}

func (r *UserRepo) VerifyEmailRegistration(ctx context.Context, req *bbsuserv1.VerifyEmailRegistration_Request) (*bbsuserv1.VerifyEmailRegistration_Reply, error) {
	_, err := r.userClient.Auth.VerifyEmailRegistration(ctx, &userv1.VerifyEmailRegistration_Request{
		Code:      req.GetCode(),
		CodeToken: req.GetCodeToken(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyEmailRegistration_Reply{}, nil
}

func (r *UserRepo) StartPhoneRegistration(ctx context.Context, req *bbsuserv1.StartPhoneRegistration_Request) (*bbsuserv1.StartPhoneRegistration_Reply, error) {
	reply, err := r.userClient.Auth.StartPhoneRegistration(ctx, &userv1.StartPhoneRegistration_Request{
		Phone:    req.GetPhone(),
		Password: req.GetPassword(),
		Name:     req.GetName(),
		Nickname: req.Nickname,
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.StartPhoneRegistration_Reply{CodeToken: reply.GetCodeToken()}, nil
}

func (r *UserRepo) VerifyPhoneRegistration(ctx context.Context, req *bbsuserv1.VerifyPhoneRegistration_Request) (*bbsuserv1.VerifyPhoneRegistration_Reply, error) {
	_, err := r.userClient.Auth.VerifyPhoneRegistration(ctx, &userv1.VerifyPhoneRegistration_Request{
		Code:      req.GetCode(),
		CodeToken: req.GetCodeToken(),
	})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.VerifyPhoneRegistration_Reply{}, nil
}

func (r *UserRepo) LoginByPassword(ctx context.Context, req *bbsuserv1.LoginByPassword_Request) (*bbsuserv1.LoginByPassword_Reply, error) {
	reply, err := r.userClient.Auth.LoginByPassword(ctx, &userv1.LoginByPassword_Request{
		Account:  req.GetAccount(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var out *bbsuserv1.Account
	if account != nil {
		out = &bbsuserv1.Account{}
		if basic := account.GetBasic(); basic != nil {
			out.Profile = &bbsuserv1.AccountProfile{
				Id:            basic.GetId(),
				Name:          basic.GetName(),
				Nickname:      basic.Nickname,
				Url:           basic.Url,
				AvatarUrl:     basic.AvatarUrl,
				Introduction:  basic.Introduction,
				Status:        bbsuserv1.AccountStatus(basic.GetStatus()),
				Mbti:          bbsuserv1.MBTI(basic.GetMbti()),
				FollowCount:   basic.FollowCount,
				FollowerCount: basic.FollowerCount,
				CreatedAt:     formatProtoTime(basic.GetCreatedAt()),
				UpdatedAt:     formatProtoTime(basic.GetUpdatedAt()),
			}
		}
		if contact := account.GetContact(); contact != nil {
			out.Contact = &bbsuserv1.AccountContact{
				UserId: contact.GetUserId(),
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &bbsuserv1.LoginByPassword_Reply{Token: reply.GetToken(), Account: out}, nil
}

func (r *UserRepo) Logout(ctx context.Context, req *bbsuserv1.Logout_Request) (*bbsuserv1.Logout_Reply, error) {
	_, err := r.userClient.Auth.Logout(forwardAuth(ctx), &userv1.Logout_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.Logout_Reply{}, nil
}

func (r *UserRepo) GetCurrentAccount(ctx context.Context, req *bbsuserv1.GetCurrentAccount_Request) (*bbsuserv1.GetCurrentAccount_Reply, error) {
	reply, err := r.userClient.Account.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentAccount_Request{})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var out *bbsuserv1.Account
	if account != nil {
		out = &bbsuserv1.Account{}
		if basic := account.GetBasic(); basic != nil {
			out.Profile = &bbsuserv1.AccountProfile{
				Id:            basic.GetId(),
				Name:          basic.GetName(),
				Nickname:      basic.Nickname,
				Url:           basic.Url,
				AvatarUrl:     basic.AvatarUrl,
				Introduction:  basic.Introduction,
				Status:        bbsuserv1.AccountStatus(basic.GetStatus()),
				Mbti:          bbsuserv1.MBTI(basic.GetMbti()),
				FollowCount:   basic.FollowCount,
				FollowerCount: basic.FollowerCount,
				CreatedAt:     formatProtoTime(basic.GetCreatedAt()),
				UpdatedAt:     formatProtoTime(basic.GetUpdatedAt()),
			}
		}
		if contact := account.GetContact(); contact != nil {
			out.Contact = &bbsuserv1.AccountContact{
				UserId: contact.GetUserId(),
				Email:  contact.Email,
				Phone:  contact.Phone,
			}
		}
	}
	return &bbsuserv1.GetCurrentAccount_Reply{Account: out}, nil
}

func (r *UserRepo) GetProfileAccount(ctx context.Context, req *bbsuserv1.GetProfileAccount_Request) (*bbsuserv1.GetProfileAccount_Reply, error) {
	reply, err := r.userClient.Account.GetBasic(ctx, &userv1.GetBasicAccount_Request{UserId: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var profile *bbsuserv1.AccountProfile
	if account != nil {
		profile = &bbsuserv1.AccountProfile{
			Id:            account.GetId(),
			Name:          account.GetName(),
			Nickname:      account.Nickname,
			Url:           account.Url,
			AvatarUrl:     account.AvatarUrl,
			Introduction:  account.Introduction,
			Status:        bbsuserv1.AccountStatus(account.GetStatus()),
			Mbti:          bbsuserv1.MBTI(account.GetMbti()),
			FollowCount:   account.FollowCount,
			FollowerCount: account.FollowerCount,
			CreatedAt:     formatProtoTime(account.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(account.GetUpdatedAt()),
		}
	}
	return &bbsuserv1.GetProfileAccount_Reply{Profile: profile}, nil
}

func (r *UserRepo) UpdateProfileAccount(ctx context.Context, req *bbsuserv1.UpdateProfileAccount_Request) (*bbsuserv1.UpdateProfileAccount_Reply, error) {
	updateReq := &userv1.UpdateProfileAccount_Request{
		AvatarUrl:    req.AvatarUrl,
		Nickname:     req.Nickname,
		Url:          req.Url,
		Introduction: req.Introduction,
	}
	if req.Mbti != nil {
		updateReq.Mbti = new(userv1.MBTI(*req.Mbti))
	}
	reply, err := r.userClient.Account.UpdateProfile(forwardAuth(ctx), updateReq)
	if err != nil {
		return nil, err
	}
	account := reply.GetAccount()
	var profile *bbsuserv1.AccountProfile
	if account != nil {
		profile = &bbsuserv1.AccountProfile{
			Id:            account.GetId(),
			Name:          account.GetName(),
			Nickname:      account.Nickname,
			Url:           account.Url,
			AvatarUrl:     account.AvatarUrl,
			Introduction:  account.Introduction,
			Status:        bbsuserv1.AccountStatus(account.GetStatus()),
			Mbti:          bbsuserv1.MBTI(account.GetMbti()),
			FollowCount:   account.FollowCount,
			FollowerCount: account.FollowerCount,
			CreatedAt:     formatProtoTime(account.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(account.GetUpdatedAt()),
		}
	}
	return &bbsuserv1.UpdateProfileAccount_Reply{Profile: profile}, nil
}

func (r *UserRepo) GetCurrentPreferences(ctx context.Context, req *bbsuserv1.GetCurrentPreferences_Request) (*bbsuserv1.GetCurrentPreferences_Reply, error) {
	reply, err := r.userClient.Preferences.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentPreferences_Request{})
	if err != nil {
		return nil, err
	}
	preferences := reply.GetPreferences()
	var out *bbsuserv1.Preference
	if preferences != nil {
		out = &bbsuserv1.Preference{
			UserId:      preferences.GetUserId(),
			Language:    preferences.Language,
			Timezone:    preferences.Timezone,
			Theme:       preferences.Theme,
			MobileTheme: preferences.MobileTheme,
		}
	}
	return &bbsuserv1.GetCurrentPreferences_Reply{Preference: out}, nil
}

func (r *UserRepo) UpdateCurrentPreferences(ctx context.Context, req *bbsuserv1.UpdateCurrentPreferences_Request) (*bbsuserv1.UpdateCurrentPreferences_Reply, error) {
	reply, err := r.userClient.Preferences.UpdateCurrent(forwardAuth(ctx), &userv1.UpdateCurrentPreferences_Request{
		Language:    req.Language,
		Timezone:    req.Timezone,
		Theme:       req.Theme,
		MobileTheme: req.MobileTheme,
	})
	if err != nil {
		return nil, err
	}
	preferences := reply.GetPreferences()
	var out *bbsuserv1.Preference
	if preferences != nil {
		out = &bbsuserv1.Preference{
			UserId:      preferences.GetUserId(),
			Language:    preferences.Language,
			Timezone:    preferences.Timezone,
			Theme:       preferences.Theme,
			MobileTheme: preferences.MobileTheme,
		}
	}
	return &bbsuserv1.UpdateCurrentPreferences_Reply{Preference: out}, nil
}

func (r *UserRepo) GetCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.GetCurrentPrivacySetting_Request) (*bbsuserv1.GetCurrentPrivacySetting_Reply, error) {
	reply, err := r.userClient.PrivacySetting.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentPrivacySetting_Request{})
	if err != nil {
		return nil, err
	}
	setting := reply.GetPrivacySetting()
	var out *bbsuserv1.PrivacySetting
	if setting != nil {
		out = &bbsuserv1.PrivacySetting{
			UserId:             setting.GetUserId(),
			PublicPoints:       setting.PublicPoints,
			PublicFollowers:    setting.PublicFollowers,
			PublicArticles:     setting.PublicArticles,
			PublicComments:     setting.PublicComments,
			PublicOnlineStatus: setting.PublicOnlineStatus,
			PublicLocation:     setting.PublicLocation,
		}
	}
	return &bbsuserv1.GetCurrentPrivacySetting_Reply{PrivacySetting: out}, nil
}

func (r *UserRepo) UpdateCurrentPrivacySetting(ctx context.Context, req *bbsuserv1.UpdateCurrentPrivacySetting_Request) (*bbsuserv1.UpdateCurrentPrivacySetting_Reply, error) {
	reply, err := r.userClient.PrivacySetting.UpdateCurrent(forwardAuth(ctx), &userv1.UpdateCurrentPrivacySetting_Request{
		PublicPoints:       req.PublicPoints,
		PublicFollowers:    req.PublicFollowers,
		PublicArticles:     req.PublicArticles,
		PublicComments:     req.PublicComments,
		PublicOnlineStatus: req.PublicOnlineStatus,
		PublicLocation:     req.PublicLocation,
	})
	if err != nil {
		return nil, err
	}
	setting := reply.GetPrivacySetting()
	var out *bbsuserv1.PrivacySetting
	if setting != nil {
		out = &bbsuserv1.PrivacySetting{
			UserId:             setting.GetUserId(),
			PublicPoints:       setting.PublicPoints,
			PublicFollowers:    setting.PublicFollowers,
			PublicArticles:     setting.PublicArticles,
			PublicComments:     setting.PublicComments,
			PublicOnlineStatus: setting.PublicOnlineStatus,
			PublicLocation:     setting.PublicLocation,
		}
	}
	return &bbsuserv1.UpdateCurrentPrivacySetting_Reply{PrivacySetting: out}, nil
}

func (r *UserRepo) GetCurrentLocation(ctx context.Context, req *bbsuserv1.GetCurrentLocation_Request) (*bbsuserv1.GetCurrentLocation_Reply, error) {
	reply, err := r.userClient.Location.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentLocation_Request{})
	if err != nil {
		return nil, err
	}
	location := reply.GetLocation()
	var out *bbsuserv1.Location
	if location != nil {
		out = &bbsuserv1.Location{
			UserId:   location.GetUserId(),
			Country:  location.Country,
			Province: location.Province,
			City:     location.City,
		}
	}
	return &bbsuserv1.GetCurrentLocation_Reply{Location: out}, nil
}

func (r *UserRepo) UpsertCurrentLocation(ctx context.Context, req *bbsuserv1.UpsertCurrentLocation_Request) (*bbsuserv1.UpsertCurrentLocation_Reply, error) {
	reply, err := r.userClient.Location.UpsertCurrent(forwardAuth(ctx), &userv1.UpsertCurrentLocation_Request{
		Country:  req.Country,
		Province: req.Province,
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	location := reply.GetLocation()
	var out *bbsuserv1.Location
	if location != nil {
		out = &bbsuserv1.Location{
			UserId:   location.GetUserId(),
			Country:  location.Country,
			Province: location.Province,
			City:     location.City,
		}
	}
	return &bbsuserv1.UpsertCurrentLocation_Reply{Location: out}, nil
}

func (r *UserRepo) Follow(ctx context.Context, req *bbsuserv1.FollowRelation_Request) (*bbsuserv1.FollowRelation_Reply, error) {
	_, err := r.userClient.Relation.Follow(forwardAuth(ctx), &userv1.FollowRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.FollowRelation_Reply{}, nil
}

func (r *UserRepo) Unfollow(ctx context.Context, req *bbsuserv1.UnfollowRelation_Request) (*bbsuserv1.UnfollowRelation_Reply, error) {
	_, err := r.userClient.Relation.Unfollow(forwardAuth(ctx), &userv1.UnfollowRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UnfollowRelation_Reply{}, nil
}

func (r *UserRepo) Block(ctx context.Context, req *bbsuserv1.BlockRelation_Request) (*bbsuserv1.BlockRelation_Reply, error) {
	_, err := r.userClient.Relation.Block(forwardAuth(ctx), &userv1.BlockRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BlockRelation_Reply{}, nil
}

func (r *UserRepo) Unblock(ctx context.Context, req *bbsuserv1.UnblockRelation_Request) (*bbsuserv1.UnblockRelation_Reply, error) {
	_, err := r.userClient.Relation.Unblock(forwardAuth(ctx), &userv1.UnblockRelation_Request{TargetId: req.GetTargetId()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.UnblockRelation_Reply{}, nil
}

func (r *UserRepo) ListFollowing(ctx context.Context, req *bbsuserv1.ListFollowingRelations_Request) (*bbsuserv1.ListFollowingRelations_Reply, error) {
	var pageReq *commonv1.PageRequest
	if req.GetPage() != nil {
		pageReq = &commonv1.PageRequest{
			Page: uint32(req.GetPage().GetCurrent()),
			Size: uint32(req.GetPage().GetPageSize()),
		}
	}
	reply, err := r.userClient.Relation.ListFollowing(forwardAuth(ctx), &userv1.ListFollowingRelations_Request{Page: pageReq})
	if err != nil {
		return nil, err
	}
	var pageReply *bbsuserv1.PageReply
	if reply.GetPage() != nil {
		pageReply = &bbsuserv1.PageReply{
			Current:  int64(reply.GetPage().GetPage()),
			PageSize: int64(reply.GetPage().GetSize()),
			Total:    int64(reply.GetPage().GetTotal()),
		}
	}
	rows := make([]*bbsuserv1.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.Relation{
			Id:        item.GetId(),
			Type:      int32(item.GetType()),
			ActorId:   item.GetActorId(),
			TargetId:  item.GetTargetId(),
			CreatedAt: formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsuserv1.ListFollowingRelations_Reply{Page: pageReply, Rows: rows}, nil
}

func (r *UserRepo) ListFollowers(ctx context.Context, req *bbsuserv1.ListFollowersRelations_Request) (*bbsuserv1.ListFollowersRelations_Reply, error) {
	var pageReq *commonv1.PageRequest
	if req.GetPage() != nil {
		pageReq = &commonv1.PageRequest{
			Page: uint32(req.GetPage().GetCurrent()),
			Size: uint32(req.GetPage().GetPageSize()),
		}
	}
	reply, err := r.userClient.Relation.ListFollowers(forwardAuth(ctx), &userv1.ListFollowersRelations_Request{Page: pageReq})
	if err != nil {
		return nil, err
	}
	var pageReply *bbsuserv1.PageReply
	if reply.GetPage() != nil {
		pageReply = &bbsuserv1.PageReply{
			Current:  int64(reply.GetPage().GetPage()),
			PageSize: int64(reply.GetPage().GetSize()),
			Total:    int64(reply.GetPage().GetTotal()),
		}
	}
	rows := make([]*bbsuserv1.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.Relation{
			Id:        item.GetId(),
			Type:      int32(item.GetType()),
			ActorId:   item.GetActorId(),
			TargetId:  item.GetTargetId(),
			CreatedAt: formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsuserv1.ListFollowersRelations_Reply{Page: pageReply, Rows: rows}, nil
}

func (r *UserRepo) ListBlocked(ctx context.Context, req *bbsuserv1.ListBlockedRelations_Request) (*bbsuserv1.ListBlockedRelations_Reply, error) {
	var pageReq *commonv1.PageRequest
	if req.GetPage() != nil {
		pageReq = &commonv1.PageRequest{
			Page: uint32(req.GetPage().GetCurrent()),
			Size: uint32(req.GetPage().GetPageSize()),
		}
	}
	reply, err := r.userClient.Relation.ListBlocked(forwardAuth(ctx), &userv1.ListBlockedRelations_Request{Page: pageReq})
	if err != nil {
		return nil, err
	}
	var pageReply *bbsuserv1.PageReply
	if reply.GetPage() != nil {
		pageReply = &bbsuserv1.PageReply{
			Current:  int64(reply.GetPage().GetPage()),
			PageSize: int64(reply.GetPage().GetSize()),
			Total:    int64(reply.GetPage().GetTotal()),
		}
	}
	rows := make([]*bbsuserv1.Relation, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		if item == nil {
			rows = append(rows, nil)
			continue
		}
		rows = append(rows, &bbsuserv1.Relation{
			Id:        item.GetId(),
			Type:      int32(item.GetType()),
			ActorId:   item.GetActorId(),
			TargetId:  item.GetTargetId(),
			CreatedAt: formatProtoTime(item.GetCreatedAt()),
			UpdatedAt: formatProtoTime(item.GetUpdatedAt()),
		})
	}
	return &bbsuserv1.ListBlockedRelations_Reply{Page: pageReply, Rows: rows}, nil
}

func (r *UserRepo) GetStatus(ctx context.Context, req *bbsuserv1.GetStatusRelation_Request) (*bbsuserv1.GetStatusRelation_Reply, error) {
	reply, err := r.userClient.Relation.BatchGetStatus(forwardAuth(ctx), &userv1.BatchGetStatusRelation_Request{TargetIds: []int64{req.GetTargetId()}})
	if err != nil {
		return nil, err
	}
	var out *bbsuserv1.RelationStatus
	if status := reply.GetStatuses()[req.GetTargetId()]; status != nil {
		out = &bbsuserv1.RelationStatus{
			TargetId:   status.GetTargetId(),
			Following:  status.GetFollowing(),
			FollowedBy: status.GetFollowedBy(),
			Blocking:   status.GetBlocking(),
			BlockedBy:  status.GetBlockedBy(),
		}
	}
	return &bbsuserv1.GetStatusRelation_Reply{Status: out}, nil
}

func (r *UserRepo) ValidateTfa(ctx context.Context, req *bbsuserv1.ValidateTfa_Request) (*bbsuserv1.ValidateTfa_Reply, error) {
	reply, err := r.userClient.Tfa.Validate(forwardAuth(ctx), &userv1.ValidateTfa_Request{Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ValidateTfa_Reply{Verified: reply.GetVerified()}, nil
}

func (r *UserRepo) BeginEnableTfa(ctx context.Context, req *bbsuserv1.BeginEnableTfa_Request) (*bbsuserv1.BeginEnableTfa_Reply, error) {
	reply, err := r.userClient.Tfa.BeginEnable(forwardAuth(ctx), &userv1.BeginEnableTfa_Request{})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.BeginEnableTfa_Reply{
		Data:        reply.GetData(),
		ContentType: reply.GetContentType(),
	}, nil
}

func (r *UserRepo) ConfirmEnableTfa(ctx context.Context, req *bbsuserv1.ConfirmEnableTfa_Request) (*bbsuserv1.ConfirmEnableTfa_Reply, error) {
	_, err := r.userClient.Tfa.ConfirmEnable(forwardAuth(ctx), &userv1.ConfirmEnableTfa_Request{Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.ConfirmEnableTfa_Reply{}, nil
}

func (r *UserRepo) DisableTfa(ctx context.Context, req *bbsuserv1.DisableTfa_Request) (*bbsuserv1.DisableTfa_Reply, error) {
	_, err := r.userClient.Tfa.Disable(forwardAuth(ctx), &userv1.DisableTfa_Request{Code: req.GetCode()})
	if err != nil {
		return nil, err
	}
	return &bbsuserv1.DisableTfa_Reply{}, nil
}

func (r *UserRepo) GetCurrentTfa(ctx context.Context, req *bbsuserv1.GetCurrentTfa_Request) (*bbsuserv1.GetCurrentTfa_Reply, error) {
	reply, err := r.userClient.Tfa.GetCurrent(forwardAuth(ctx), &userv1.GetCurrentTfa_Request{})
	if err != nil {
		return nil, err
	}
	tfa := reply.GetTfa()
	var out *bbsuserv1.Tfa
	if tfa != nil {
		out = &bbsuserv1.Tfa{
			UserId:     tfa.GetUserId(),
			Enable:     tfa.GetEnable(),
			EnableTime: formatProtoTime(tfa.GetEnableTime()),
		}
	}
	return &bbsuserv1.GetCurrentTfa_Reply{Tfa: out}, nil
}
